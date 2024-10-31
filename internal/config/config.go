package config

import "path/filepath"

type IndexOpts struct {
	ModuleRoot    string
	ModuleVersion string

	// Path for the current module we are indexing. Same as packages.Package.Module.Path
	ModulePath string

	// Go version. Used for linking to the Go standard library
	GoStdlibVersion string

	// Whether we should emit implementations
	SkipImplementations bool
	SkipTests           bool

	IsIndexingStdlib bool

	// Package patterns to index
	PackagePatterns []string
}

func New(ModuleRoot, ModuleVersion, ModulePath, GoStdlibVersion string, IsIndexingStdlib bool, SkipImplementations bool, SkipTests bool, PackagePatterns []string) IndexOpts {
	ModuleRoot, err := filepath.Abs(ModuleRoot)
	if err != nil {
		panic(err)
	}

	return IndexOpts{
		ModuleRoot:          ModuleRoot,
		ModuleVersion:       ModuleVersion,
		ModulePath:          ModulePath,
		GoStdlibVersion:     GoStdlibVersion,
		SkipImplementations: SkipImplementations,
		SkipTests:           SkipTests,
		IsIndexingStdlib:    IsIndexingStdlib,
		PackagePatterns:     PackagePatterns,
	}
}
