package index

import (
	"go/ast"
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

// findPackageDoc returns the doc comment text for the package, picking the
// best file when multiple files have doc comments. Returns "" if no file
// has a package doc comment.
func findPackageDoc(pkg *packages.Package) string {
	var filesWithDocs []*ast.File
	for _, f := range pkg.Syntax {
		if f.Doc != nil {
			filesWithDocs = append(filesWithDocs, f)
		}
	}

	candidates := filterBasedOnTestFiles(pkg, filesWithDocs)
	if len(candidates) == 0 {
		return ""
	}

	// Prefer doc.go, then the file matching the package name, then pick the
	// first candidate.
	for _, f := range candidates {
		if path.Base(pkg.Fset.Position(f.Pos()).Filename) == "doc.go" {
			return f.Doc.Text()
		}
	}
	for _, f := range candidates {
		if fileNameWithoutExtension(pkg.Fset.Position(f.Pos()).Filename) == pkg.Name {
			return f.Doc.Text()
		}
	}
	return candidates[0].Doc.Text()
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
