package index

import (
	"go/ast"
	"math"
	"path"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/sourcegraph/scip-go/internal/handler"
	"golang.org/x/tools/go/packages"
)

func findBestPackageDefinitionPath(pkg *packages.Package) (*ast.File, error) {
	if pkg.PkgPath == "builtin" {
		return nil, nil
	}

	// Unsafe is special case for builtin
	if pkg.PkgPath == "unsafe" {
		return nil, nil
	}

	if len(pkg.Syntax) == 0 {
		handler.Println("Missing |", pkg.ID, pkg.Module.Path)
		return nil, nil
	}

	files := []*ast.File{}
	filesWithDocs := []*ast.File{}
	for _, f := range pkg.Syntax {
		// pos := pkg.Fset.Position(f.Pos())

		files = append(files, f)
		if f.Doc != nil {
			filesWithDocs = append(filesWithDocs, f)
		}
	}

	// The idiomatic way is to _only_ have one .go file per package that has a docstring
	// for the package. This should generally return here.
	if len(filesWithDocs) == 1 {
		return filesWithDocs[0], nil
	}

	// If we for some reason have more than one .go file per package that has a docstring,
	// only consider returning paths that contain the docstring (instead of any of the possible
	// paths).
	if len(filesWithDocs) > 1 {
		files = filesWithDocs
	}

	// Try to only pick non _test files for non _test packages and vice versa.
	files = filterBasedOnTestFiles(pkg, files)

	// Find the best remaining path.
	// Chooses:
	//     1. doc.go
	//     2. exact match
	//     3. computes levenshtein and picks best score
	var bestFile *ast.File

	minDistance := math.MaxInt32
	for _, f := range files {
		fPath := pkg.Fset.Position(f.Pos()).Filename
		fileName := fileNameWithoutExtension(fPath)

		if "doc.go" == path.Base(fPath) {
			return f, nil
		}

		if pkg.Name == fileName {
			return f, nil
		}

		distance := levenshtein.ComputeDistance(pkg.Name, fileName)
		if distance < minDistance {
			minDistance = distance
			bestFile = f
		}
	}

	return bestFile, nil
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, path.Ext(fileName))
}

func filterBasedOnTestFiles(pkg *packages.Package, files []*ast.File) []*ast.File {
	packageNameEndsWithTest := strings.HasSuffix(pkg.Name, "_test")

	preferredFiles := []*ast.File{}
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
