package main

import (
	"encoding/json"
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
	"github.com/sourcegraph/scip-go/internal/index"
	"github.com/sourcegraph/scip-go/internal/modules"
	"github.com/sourcegraph/scip-go/internal/output"
	"golang.org/x/tools/go/packages"
	"google.golang.org/protobuf/proto"
)

type SharedFlags struct {
	ModuleRoot          string   `help:"Specifies the directory containing the go.mod file." default:"${module_root}"`
	RepositoryRemote    string   `help:"Specifies the canonical name of the repository remote." default:"${repository_remote}"`
	ModulePath          string   `help:"Overrides the module path inferred from go.mod."`
	ModuleVersion       string   `help:"Specifies the version of the module defined by module-root." default:"${module_version}"`
	GoVersion           string   `help:"Specifies the version of the Go standard library to link to. Format: 'go1.XX'" default:"${go_version}"`
	Quiet               bool     `help:"Do not output to stdout or stderr." short:"q"`
	Verbose             int      `help:"Output debug logs." short:"V" type:"counter"`
	SkipImplementations bool     `help:"Skip implementations. Use to skip generating implementations"`
	SkipTests           bool     `help:"Skip compiling tests. Will not generate scip indexes over your or your dependencies tests"`
	PackagePatterns     []string `arg:"" optional:"" help:"Package patterns to index. Default: './...' which indexes all packages in the current directory recursively. For the full syntax of allowed package patterns, see https://pkg.go.dev/cmd/go#hdr-Package_lists_and_patterns" default:"./..."`
}

type IndexCmd struct {
	SharedFlags
	Output  string `help:"The output file." short:"o" default:"index.scip"`
	Profile int    `help:"Turn on debug profiling. This will reduce performance. Do not turn on unless debugging. Set to number of milliseconds per sample"`
}

type PackagesCmd struct {
	SharedFlags
}

type MissingCmd struct {
	SharedFlags
}

type CLI struct {
	Index    IndexCmd         `cmd:"" default:"withargs" help:"Index Go source code and emit an SCIP index."`
	Packages PackagesCmd      `cmd:"" help:"List current and dependency packages."`
	Missing  MissingCmd       `cmd:"" help:"List missing documents."`
	Version  kong.VersionFlag `help:"Show version." short:"v"`
}

func main() {
	ctx, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := ctx.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func makeOptions(shared *SharedFlags) (config.IndexOpts, error) {
	moduleRoot, err := filepath.Abs(shared.ModuleRoot)
	if err != nil {
		return config.IndexOpts{}, fmt.Errorf("get abspath of module root: %v", err)
	}

	output.SetOutputOptions(getLogLevel(shared.Quiet, shared.Verbose))

	modulePath, isStdLib, err := modules.ModuleName(moduleRoot, shared.RepositoryRemote, shared.ModulePath)
	if err != nil {
		return config.IndexOpts{}, err
	}

	slog.Info("Go standard library version: ", "version", shared.GoVersion)
	slog.Info("Resolved module name: ", "module", modulePath)
	if isStdLib {
		slog.Info("Resolved as stdlib: true")
	}
	if shared.SkipImplementations {
		slog.Info("Skipping implementations")
	}
	if shared.SkipTests {
		slog.Info("Skipping tests")
	}

	return config.New(moduleRoot, shared.ModuleVersion, modulePath, shared.GoVersion, isStdLib, shared.SkipImplementations, shared.SkipTests, shared.PackagePatterns), nil
}

func (cmd *IndexCmd) Run() error {
	if cmd.Profile > 0 {
		runtime.MemProfileRate = cmd.Profile
		f, err := os.Create("mem.pprof")
		if err != nil {
			return fmt.Errorf("could not create memory profile: %w", err)
		}
		defer func() {
			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				slog.Error("could not write memory profile", "error", err)
			}
			if err := f.Close(); err != nil {
				slog.Error("could not close memory profile", "error", err)
			}
		}()
	}

	options, err := makeOptions(&cmd.SharedFlags)
	if err != nil {
		return err
	}

	file, err := os.Create(cmd.Output)
	if err != nil {
		return fmt.Errorf("failed to create scip index file %q: %w", cmd.Output, err)
	}
	defer file.Close()

	var fileMutex sync.Mutex
	writer := func(msg proto.Message) error {
		index := &scip.Index{}

		switch msg := msg.(type) {
		case *scip.Metadata:
			index.Metadata = msg
		case *scip.Document:
			index.Documents = append(index.Documents, msg)
		case *scip.SymbolInformation:
			index.ExternalSymbols = append(index.ExternalSymbols, msg)
		default:
			return fmt.Errorf("invalid msg type %T", msg)
		}

		b, err := proto.Marshal(index)
		if err != nil {
			return fmt.Errorf("failed to marshal SCIP index to Protobuf: %w", err)
		}

		// Lock for writing to the file, to make sure that we don't race to
		// write things (serializing can be done before locking)
		fileMutex.Lock()
		defer fileMutex.Unlock()

		if _, err := file.Write(b); err != nil {
			return fmt.Errorf("failed to write to scip index file: %w", err)
		}

		return nil
	}

	removeOutFileIfPresent := func() {
		if fileInfo, err := os.Stat(cmd.Output); err == nil && fileInfo.Mode().IsRegular() {
			os.RemoveAll(cmd.Output)
		}
	}

	if err = index.Index(writer, options); err != nil {
		removeOutFileIfPresent()
		return err
	}

	return nil
}

func (cmd *PackagesCmd) Run() error {
	options, err := makeOptions(&cmd.SharedFlags)
	if err != nil {
		return err
	}

	current, deps, err := index.GetPackages(options)
	if err != nil {
		return err
	}

	fmt.Println("Current packages")
	for _, pkgID := range current {
		fmt.Println(string(pkgID))
	}

	fmt.Println("Dependency packages")
	for _, pkgID := range deps {
		fmt.Println(string(pkgID))
	}
	return nil
}

func (cmd *MissingCmd) Run() error {
	options, err := makeOptions(&cmd.SharedFlags)
	if err != nil {
		return err
	}

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

func parseArgs(args []string) (*kong.Context, error) {
	cli := &CLI{}

	parser, err := kong.New(cli,
		kong.Name("scip-go"),
		kong.Description("scip-go is an SCIP indexer for Go."),
		kong.DefaultEnvars(""),
		kong.Vars{
			"version":           index.ScipGoVersion,
			"module_root":       defaultModuleRoot(),
			"repository_remote": defaultRepositoryRemote(),
			"module_version":    defaultModuleVersion(),
			"go_version":        defaultGoVersion(),
		},
		kong.UsageOnError(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %v", err)
	}

	ctx, err := parser.Parse(args)
	if err != nil {
		return nil, fmt.Errorf("failed to parse args: %v", err)
	}

	return ctx, nil
}

//
// Defaults

var defaultModuleRoot = sync.OnceValue(func() string {
	return searchForGoMod(wd(), toplevel())
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

var logLevels = []slog.Level{
	slog.LevelWarn,  // default
	slog.LevelInfo,  // -V
	slog.LevelDebug, // -VV or more
}

func getLogLevel(noOutput bool, verbosity int) slog.Level {
	if noOutput {
		return slog.LevelError
	}

	if verbosity >= len(logLevels) {
		verbosity = len(logLevels) - 1
	}

	return logLevels[verbosity]
}
