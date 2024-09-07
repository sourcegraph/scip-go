package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/alecthomas/kingpin"
	"github.com/charmbracelet/log"
	"github.com/pkg/profile"
	"github.com/sourcegraph/scip-go/internal/command"
	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/git"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/index"
	"github.com/sourcegraph/scip-go/internal/modules"
	"github.com/sourcegraph/scip-go/internal/output"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
	"google.golang.org/protobuf/proto"
)

var app = kingpin.New(
	"scip-go",
	"scip-go is an SCIP indexer for Go.",
).Version(index.ScipGoVersion)

var (
	outFile          string
	projectRoot      string
	moduleRoot       string
	repositoryRoot   string
	repositoryRemote string
	moduleVersion    string
	moduleName       string
	goVersion        string
	verbosity        int
	noOutput         bool
	animation        bool
	devMode          bool

	// fUnNy cOmMaNd
	scipCommand string

	// TODO: We should consider if we can avoid doing this in this iteration of scip-go
	// depBatchSize          int
	skipImplementations bool
	skipTests           bool

	// Debugging flag to turn on profiling
	profileRate int
)

func init() {
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('V')

	// Outfile options
	app.Flag("output", "The output file.").Short('o').Default("index.scip").StringVar(&outFile)

	// Path options (inferred by presence of go.mod; git)
	app.Flag("project-root", "Specifies the directory to index.").Default(".").StringVar(&projectRoot)
	app.Flag("module-root", "Specifies the directory containing the go.mod file.").Default(defaultModuleRoot()).StringVar(&moduleRoot)
	app.Flag("repository-root", "Specifies the top-level directory of the git repository.").Default(defaultRepositoryRoot()).StringVar(&repositoryRoot)

	// Repository remote and tag options (inferred by git)
	app.Flag("repository-remote", "Specifies the canonical name of the repository remote.").Default(defaultRepositoryRemote()).StringVar(&repositoryRemote)
	app.Flag("module-name", "Specifies the name of the module defined by module-root.").StringVar(&moduleName)
	app.Flag("module-version", "Specifies the version of the module defined by module-root.").Default(defaultModuleVersion()).StringVar(&moduleVersion)
	app.Flag("go-version", "Specifies the version of the Go standard library to link to. Format: 'go1.XX'").Default(defaultGoVersion()).StringVar(&goVersion)

	// Verbosity options
	app.Flag("quiet", "Do not output to stdout or stderr.").Short('q').Default("false").BoolVar(&noOutput)
	app.Flag("verbose", "Output debug logs.").Short('v').CounterVar(&verbosity)
	app.Flag("animation", "Show animated output.").Default("false").BoolVar(&animation)
	app.Flag("dev", "Enable development mode.").Default("false").BoolVar(&devMode)

	// app.Flag("dep-batch-size", "How many dependencies to load at once to limit memory usage (e.g. 100). 0 means load all at once.").Default("0").IntVar(&depBatchSize)
	app.Flag("skip-implementations", "Skip implementations. Use to skip generating implementations").Default("false").BoolVar(&skipImplementations)
	app.Flag("skip-tests", "Skip compiling tests. Will not generate scip indexes over your or your dependencies tests").Default("false").BoolVar(&skipTests)

	app.Flag("command", "Optionally specifies a command to run. Defaults to 'index'").Default("index").StringVar(&scipCommand)

	app.Flag("profile", "Turn on debug profiling. This will reduce performance. Do not turn on unless debugging. Set to number of milliseconds per sample").Default("0").IntVar(&profileRate)
}

func main() {
	if err := mainErr(); err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("error: %v\n", err))
		os.Exit(1)
	}
}

func mainErr() error {
	if err := parseArgs(os.Args[1:]); err != nil {
		return err
	}

	if profileRate > 0 {
		p := profile.MemProfileRate(profileRate)
		defer profile.Start(p).Stop()
	}

	handler.SetDev(devMode)

	output.SetOutputOptions(getVerbosity(), animation)
	// The default formatting also prints the date, which is generally not needed.
	log.SetTimeFormat("15:04:05")
	log.SetStyles(func() *log.Styles {
		// The default styles will print 'DEBU' and 'ERRO' to line
		// up with 'INFO' and 'WARN', instead of 'DEBUG' and 'ERROR'
		s := log.DefaultStyles()
		for lvl, style := range s.Levels {
			s.Levels[lvl] = style.UnsetMaxWidth()
		}
		return s
	}())

	modulePath, isStdLib, err := modules.ModuleName(moduleRoot, repositoryRemote, moduleName)

	log.Info("Go standard library version: ", "version", goVersion)
	log.Info("Resolved module name: ", "module", modulePath)
	if isStdLib {
		log.Info("Resolved as stdlib: true")
	}
	if skipImplementations {
		log.Info("Skipping implementations")
	}
	if skipTests {
		log.Info("Skipping tests")
	}

	options := config.New(moduleRoot, moduleVersion, modulePath, goVersion, isStdLib, skipImplementations, skipTests)

	if strings.HasPrefix(scipCommand, "list-packages") {
		var filter string
		if strings.Contains(scipCommand, ":") {
			filter = strings.Split(scipCommand, ":")[1]
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

	if scipCommand == "list-missing" {
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

	file, err := os.Create(outFile)
	if err != nil {
		log.Fatal("Failed to create scip index file", "path", outFile)
		return err
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
			log.Fatal("Failed to marshal SCIP index to Protobuf")
		}

		// Lock for writing to the file, to make sure that we don't race to
		// write things (serializing can be done before locking)
		fileMutex.Lock()
		defer fileMutex.Unlock()

		// Serialize to file. Items can now be discarded
		if _, err := file.Write(b); err != nil {
			log.Fatal("Failed to write to scip index file")
		}
	}

	removeOutFileIfPresent := func() {
		if fileInfo, err := os.Stat(outFile); err == nil && fileInfo.Mode().IsRegular() {
			os.RemoveAll(outFile)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			removeOutFileIfPresent()
		}
	}()

	if err := index.Index(writer, options); err != nil {
		removeOutFileIfPresent()
		return err
	}

	return nil
}

func parseArgs(args []string) (err error) {
	if _, err := app.Parse(args); err != nil {
		return fmt.Errorf("failed to parse args: %v", err)
	}

	sanitizers := []func() error{sanitizeProjectRoot, sanitizeModuleRoot, sanitizeRepositoryRoot}
	validators := []func() error{validatePaths}

	for _, f := range append(sanitizers, validators...) {
		if err := f(); err != nil {
			return fmt.Errorf("failed to parse args: %v", err)
		}
	}

	return nil
}

//
// Sanitizers

func sanitizeProjectRoot() (err error) {
	projectRoot, err = filepath.Abs(projectRoot)
	if err != nil {
		return fmt.Errorf("get abspath of project root: %v", err)
	}

	return nil
}

func sanitizeModuleRoot() (err error) {
	moduleRoot, err = filepath.Abs(moduleRoot)
	if err != nil {
		return fmt.Errorf("get abspath of module root: %v", err)
	}

	return nil
}

func sanitizeRepositoryRoot() (err error) {
	repositoryRoot, err = filepath.Abs(repositoryRoot)
	if err != nil {
		return fmt.Errorf("get abspath of repository root: %v", err)
	}

	return nil
}

//
// Validators

func validatePaths() error {
	if !strings.HasPrefix(projectRoot, repositoryRoot) {
		return errors.New("project root is outside the repository")
	}

	if !strings.HasPrefix(moduleRoot, repositoryRoot) {
		return errors.New("module root is outside the repository")
	}

	return nil
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
	modOutput, err := command.Run(moduleRoot, "go", "list", "-mod=readonly", "-m", "-json")
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

func getVerbosity() output.Verbosity {
	if noOutput {
		return output.NoOutput
	}

	if verbosity >= len(verbosityLevels) {
		verbosity = len(verbosityLevels) - 1
	}

	return verbosityLevels[verbosity]
}
