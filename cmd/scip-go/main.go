package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/git"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/index"
	"github.com/sourcegraph/scip-go/internal/modules"
	"github.com/sourcegraph/scip-go/internal/output"
	"google.golang.org/protobuf/proto"
)

var app = kingpin.New(
	"scip-go",
	"scip-go is an SCIP indexer for Go.",
).Version("0.1")

var (
	outFile          string
	projectRoot      string
	moduleRoot       string
	repositoryRoot   string
	repositoryRemote string
	moduleVersion    string
	moduleName       string
	verbosity        int
	noOutput         bool
	animation        bool
	devMode          bool

	scipCommand string

	// TODO: We should consider if we can avoid doing this in this iteration of scip-go
	// depBatchSize          int
)

func init() {
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('V')

	// Outfile options
	app.Flag("output", "The output file.").Short('o').Default("index.scip").StringVar(&outFile)

	// Path options (inferred by presence of go.mod; git)
	app.Flag("project-root", "Specifies the directory to index.").Default(".").StringVar(&projectRoot)
	app.Flag("module-root", "Specifies the directory containing the go.mod file.").Default(defaultModuleRoot.Value()).StringVar(&moduleRoot)
	app.Flag("repository-root", "Specifies the top-level directory of the git repository.").Default(defaultRepositoryRoot.Value()).StringVar(&repositoryRoot)

	// Repository remote and tag options (inferred by git)
	app.Flag("repository-remote", "Specifies the canonical name of the repository remote.").Default(defaultRepositoryRemote.Value()).StringVar(&repositoryRemote)
	app.Flag("module-name", "Specifies the name of the module defined by module-root.").StringVar(&moduleName)
	app.Flag("module-version", "Specifies the version of the module defined by module-root.").Default(defaultModuleVersion.Value()).StringVar(&moduleVersion)

	// Verbosity options
	app.Flag("quiet", "Do not output to stdout or stderr.").Short('q').Default("false").BoolVar(&noOutput)
	app.Flag("verbose", "Output debug logs.").Short('v').CounterVar(&verbosity)
	app.Flag("animation", "Do not animate output.").Default("false").BoolVar(&animation)
	app.Flag("dev", "Enable development mode.").Default("false").BoolVar(&devMode)

	// app.Flag("dep-batch-size", "How many dependencies to load at once to limit memory usage (e.g. 100). 0 means load all at once.").Default("0").IntVar(&depBatchSize)

	app.Flag("command", "Optionally specifies a command to run. Defaults to 'index'").Default("index").StringVar(&scipCommand)
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

	handler.SetDev(devMode)

	output.SetOutputOptions(getVerbosity(), animation)
	output.Println("scip-go")

	modulePath, isStd, err := modules.ModuleName(moduleRoot, repositoryRemote, moduleName)
	if isStd {
		panic("TODO: support stdlib. Check old lsif-go status")
	}

	options := config.New(moduleRoot, moduleVersion, modulePath)

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

	index, err := index.Index(options)
	if err != nil {
		return err
	}

	b, err := proto.Marshal(index)
	if err != nil {
		return err
	}

	return os.WriteFile(outFile, b, 0644)
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

var defaultModuleRoot = newCachedString(func() string {
	return searchForGoMod(wd.Value(), toplevel.Value())
})

var defaultRepositoryRoot = newCachedString(func() string {
	return rel(toplevel.Value())
})

var defaultRepositoryRemote = newCachedString(func() string {
	if repo, err := git.InferRepo(defaultModuleRoot.Value()); err == nil {
		return repo
	}

	return ""
})

var defaultModuleVersion = newCachedString(func() string {
	if version, err := git.InferModuleVersion(defaultModuleRoot.Value()); err == nil {
		return version
	}

	return ""
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
