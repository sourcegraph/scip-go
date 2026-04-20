package config

type IndexOpts struct {
	ModuleRoot    string
	ModuleVersion string

	// Path for the current module we are indexing. Same as packages.Package.Module.Path
	ModulePath string

	IsGoPackagesDriverSet bool

	// Go version. Used for linking to the Go standard library
	GoStdlibVersion string

	// Whether we should emit implementations
	SkipImplementations bool
	SkipTests           bool

	IsIndexingStdlib bool

	// Package patterns to index
	PackagePatterns []string

	// Arguments passed to the CLI
	Arguments []string
}
