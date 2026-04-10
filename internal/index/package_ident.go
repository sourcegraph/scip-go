package index

import (
	"go/ast"
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

// findPackageDocFile picks the file whose doc comment should represent
// the package's documentation in the SCIP index. Returns nil if no file
// has a package doc comment.
func findPackageDocFile(pkg *packages.Package) *ast.File {
	var filesWithDocs []*ast.File
	for _, f := range pkg.Syntax {
		if f.Doc != nil {
			filesWithDocs = append(filesWithDocs, f)
		}
	}

	if len(filesWithDocs) == 1 {
		return filesWithDocs[0]
	}

	if len(filesWithDocs) == 0 {
		return nil
	}

	// Multiple files have doc comments. Prefer doc.go, then the file matching
	// the package name, then pick the first candidate.
	candidates := filterBasedOnTestFiles(pkg, filesWithDocs)
	for _, f := range candidates {
		if path.Base(pkg.Fset.Position(f.Pos()).Filename) == "doc.go" {
			return f
		}
	}
	for _, f := range candidates {
		if fileNameWithoutExtension(pkg.Fset.Position(f.Pos()).Filename) == pkg.Name {
			return f
		}
	}

	return candidates[0]
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, path.Ext(fileName))
}

func filterBasedOnTestFiles(pkg *packages.Package, files []*ast.File) []*ast.File {
	packageNameEndsWithTest := strings.HasSuffix(pkg.Name, "_test")

	var preferredFiles []*ast.File
	for _, f := range files {
		fPath := pkg.Fset.Position(f.Pos())
		if packageNameEndsWithTest == strings.HasSuffix(fPath.Filename, "_test.go") {
			preferredFiles = append(preferredFiles, f)
		}
	}

	if len(preferredFiles) > 0 {
		return preferredFiles
	}

	return files
}
