package config

import "path/filepath"

type IndexOpts struct {
	ModuleRoot    string
	ModuleVersion string

	// Path for the current module we are indexing. Same as packages.Package.Module.Path
	ModulePath string

	// Go version. Used for linking to the Go standard library
	GoStdlibVersion string

	IsIndexingStdlib bool
}

func New(moduleRoot, moduleVersion, modulePath, goVersion string, isIndexingStdlib bool) IndexOpts {
	moduleRoot, err := filepath.Abs(moduleRoot)
	if err != nil {
		panic(err)
	}

	return IndexOpts{
		moduleRoot,
		moduleVersion,
		modulePath,
		goVersion,
		isIndexingStdlib,
	}
}
