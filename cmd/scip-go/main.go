package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"

	"github.com/alecthomas/kong"
	"github.com/scip-code/scip/bindings/go/scip"
	"github.com/sourcegraph/scip-go/internal/command"
	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/git"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/index"
	"github.com/sourcegraph/scip-go/internal/modules"
	"github.com/sourcegraph/scip-go/internal/output"
	"golang.org/x/tools/go/packages"
	"google.golang.org/protobuf/proto"
)

type CLI struct {
	Output              string           `help:"The output file." short:"o" default:"index.scip"`
	ProjectRoot         string           `help:"Specifies the directory to index." default:"."`
	ModuleRoot          string           `help:"Specifies the directory containing the go.mod file."`
	RepositoryRoot      string           `help:"Specifies the top-level directory of the git repository."`
	RepositoryRemote    string           `help:"Specifies the canonical name of the repository remote."`
	ModuleName          string           `help:"Specifies the name of the module defined by module-root."`
	ModuleVersion       string           `help:"Specifies the version of the module defined by module-root."`
	GoVersion           string           `help:"Specifies the version of the Go standard library to link to. Format: 'go1.XX'" name:"go-version"`
	Quiet               bool             `help:"Do not output to stdout or stderr." short:"q"`
	Verbose             int              `help:"Output debug logs." short:"v" type:"counter"`
	Dev                 bool             `help:"Enable development mode."`
	SkipImplementations bool             `help:"Skip implementations. Use to skip generating implementations" name:"skip-implementations"`
	SkipTests           bool             `help:"Skip compiling tests. Will not generate scip indexes over your or your dependencies tests" name:"skip-tests"`
	Command             string           `help:"Optionally specifies a command to run. Defaults to 'index'" default:"index"`
	Profile             int              `help:"Turn on debug profiling. This will reduce performance. Do not turn on unless debugging. Set to number of milliseconds per sample"`
	PackagePatterns     []string         `arg:"" optional:"" help:"Package patterns to index. Default: './...' which indexes all packages in the current directory recursively. For the full syntax of allowed package patterns, see https://pkg.go.dev/cmd/go#hdr-Package_lists_and_patterns" default:"./..."`
	Version             kong.VersionFlag `help:"Show version." short:"V"`
}

func main() {
	if err := mainErr(); err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("error: %v\n", err))
		os.Exit(1)
	}
}

func mainErr() (err error) {
	cli, err := parseArgs(os.Args[1:])
	if err != nil {
		return err
	}

	if cli.Profile > 0 {
		runtime.MemProfileRate = cli.Profile
		f, err := os.Create("mem.pprof")
		if err != nil {
			return fmt.Errorf("could not create memory profile: %w", err)
		}
		defer func() {
			runtime.GC()
			pprof.WriteHeapProfile(f)
			f.Close()
		}()
	}

	handler.SetDev(cli.Dev)

	output.SetOutputOptions(getVerbosity(cli.Quiet, cli.Verbose))

	modulePath, isStdLib, err := modules.ModuleName(cli.ModuleRoot, cli.RepositoryRemote, cli.ModuleName)

	slog.Info("Go standard library version: ", "version", cli.GoVersion)
	slog.Info("Resolved module name: ", "module", modulePath)
	if isStdLib {
		slog.Info("Resolved as stdlib: true")
	}
	if cli.SkipImplementations {
		slog.Info("Skipping implementations")
	}
	if cli.SkipTests {
		slog.Info("Skipping tests")
	}

	options := config.New(cli.ModuleRoot, cli.ModuleVersion, modulePath, cli.GoVersion, isStdLib, cli.SkipImplementations, cli.SkipTests, cli.PackagePatterns)

	if strings.HasPrefix(cli.Command, "list-packages") {
		var filter string
		if strings.Contains(cli.Command, ":") {
			filter = strings.Split(cli.Command, ":")[1]
		}

		current, deps, err := index.GetPackages(options)
		if err != nil {
			return err
		}

		fmt.Println("Current packages")
		for _, pkgID := range current {
			pkg := string(pkgID)
			if filter == "" || strings.Contains(pkg, filter) {
				fmt.Println(pkg)
			}
		}

		fmt.Println("Dependency packages")
		for _, pkgID := range deps {
			pkg := string(pkgID)
			if filter == "" || strings.Contains(pkg, filter) {
				fmt.Println(pkg)
			}
		}
		return nil
	}

	if cli.Command == "list-missing" {
		missing, err := index.ListMissing(options)
		if err != nil {
			return err
		}

		if len(missing) == 0 {
			fmt.Println("No missing documents")
		} else {
			fmt.Println("Missing documents:")
			for _, m := range missing {
				fmt.Println(m)
			}
		}
		return nil
	}

	file, err := os.Create(cli.Output)
	if err != nil {
		panic(fmt.Sprintf("failed to create scip index file %q: %v", cli.Output, err))
	}
	defer file.Close()

	var fileMutex sync.Mutex
	writer := func(msg proto.Message) {
		index := &scip.Index{}

		switch msg := msg.(type) {
		case *scip.Metadata:
			index.Metadata = msg
		case *scip.Document:
			index.Documents = append(index.Documents, msg)
		case *scip.SymbolInformation:
			index.ExternalSymbols = append(index.ExternalSymbols, msg)
		default:
			panic("invalid msg type")
		}

		b, err := proto.Marshal(index)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal SCIP index to Protobuf: %v", err))
		}

		// Lock for writing to the file, to make sure that we don't race to
		// write things (serializing can be done before locking)
		fileMutex.Lock()
		defer fileMutex.Unlock()

		// Serialize to file. Items can now be discarded
		if _, err := file.Write(b); err != nil {
			panic(fmt.Sprintf("failed to write to scip index file: %v", err))
		}
	}

	removeOutFileIfPresent := func() {
		if fileInfo, err := os.Stat(cli.Output); err == nil && fileInfo.Mode().IsRegular() {
			os.RemoveAll(cli.Output)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			removeOutFileIfPresent()
			err = fmt.Errorf("panic during indexing: %v", r)
		}
	}()

	if err = index.Index(writer, options); err != nil {
		removeOutFileIfPresent()
		return err
	}

	return nil
}

func parseArgs(args []string) (*CLI, error) {
	cli := &CLI{}

	// Compute defaults that depend on the environment
	moduleRootDefault := defaultModuleRoot()
	repoRootDefault := defaultRepositoryRoot()
	repoRemoteDefault := defaultRepositoryRemote()
	moduleVersionDefault := defaultModuleVersion()
	goVersionDefault := defaultGoVersion()

	parser, err := kong.New(cli,
		kong.Name("scip-go"),
		kong.Description("scip-go is an SCIP indexer for Go."),
		kong.DefaultEnvars(""),
		kong.Vars{
			"version": index.ScipGoVersion,
		},
		kong.UsageOnError(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %v", err)
	}

	if _, err := parser.Parse(args); err != nil {
		return nil, fmt.Errorf("failed to parse args: %v", err)
	}

	// Apply dynamic defaults for flags that weren't explicitly set
	if cli.ModuleRoot == "" {
		cli.ModuleRoot = moduleRootDefault
	}
	if cli.RepositoryRoot == "" {
		cli.RepositoryRoot = repoRootDefault
	}
	if cli.RepositoryRemote == "" {
		cli.RepositoryRemote = repoRemoteDefault
	}
	if cli.ModuleVersion == "" {
		cli.ModuleVersion = moduleVersionDefault
	}
	if cli.GoVersion == "" {
		cli.GoVersion = goVersionDefault
	}

	// Sanitize paths
	cli.ProjectRoot, err = filepath.Abs(cli.ProjectRoot)
	if err != nil {
		return nil, fmt.Errorf("get abspath of project root: %v", err)
	}
	cli.ModuleRoot, err = filepath.Abs(cli.ModuleRoot)
	if err != nil {
		return nil, fmt.Errorf("get abspath of module root: %v", err)
	}
	cli.RepositoryRoot, err = filepath.Abs(cli.RepositoryRoot)
	if err != nil {
		return nil, fmt.Errorf("get abspath of repository root: %v", err)
	}

	// Validate
	if !strings.HasPrefix(cli.ProjectRoot, cli.RepositoryRoot) {
		return nil, errors.New("project root is outside the repository")
	}
	if !strings.HasPrefix(cli.ModuleRoot, cli.RepositoryRoot) {
		return nil, errors.New("module root is outside the repository")
	}

	return cli, nil
}

//
// Defaults

var defaultModuleRoot = sync.OnceValue(func() string {
	return searchForGoMod(wd(), toplevel())
})

var defaultRepositoryRoot = sync.OnceValue(func() string {
	return rel(toplevel())
})

var defaultRepositoryRemote = sync.OnceValue(func() string {
	if repo, err := git.InferRepo(defaultModuleRoot()); err == nil {
		return repo
	}

	return ""
})

var defaultModuleVersion = sync.OnceValue(func() string {
	if version, err := git.InferModuleVersion(defaultModuleRoot()); err == nil {
		return version
	}

	return ""
})

var defaultGoVersion = sync.OnceValue(func() string {
	modOutput, err := command.Run(defaultModuleRoot(), "go", "list", "-mod=readonly", "-m", "-json")
	if err != nil {
		return ""
	}

	var thisPackage *packages.Module
	if err := json.NewDecoder(strings.NewReader(modOutput)).Decode(&thisPackage); err != nil {
		return ""
	}

	return "go" + thisPackage.GoVersion
})

var verbosityLevels = map[int]output.Verbosity{
	0: output.DefaultOutput,
	1: output.VerboseOutput,
	2: output.VeryVerboseOutput,
	3: output.VeryVeryVerboseOutput,
}

func getVerbosity(noOutput bool, verbosity int) output.Verbosity {
	if noOutput {
		return output.NoOutput
	}

	if verbosity >= len(verbosityLevels) {
		verbosity = len(verbosityLevels) - 1
	}

	return verbosityLevels[verbosity]
}
