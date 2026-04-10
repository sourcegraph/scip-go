package index

import (
	"go/ast"
	"path"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

// findPackageDocs returns all doc comment texts for the package, sorted by
// relevance: doc.go first, then the file matching the package name, then
// other matching files, and finally test/non-test mismatched files last.
// Returns nil if no file has a package doc comment.
func findPackageDocs(pkg *packages.Package) []string {
	var filesWithDocs []*ast.File
	for _, f := range pkg.Syntax {
		if f.Doc != nil {
			filesWithDocs = append(filesWithDocs, f)
		}
	}

	sort.SliceStable(filesWithDocs, func(i, j int) bool {
		fi := pkg.Fset.Position(filesWithDocs[i].Pos()).Filename
		fj := pkg.Fset.Position(filesWithDocs[j].Pos()).Filename
		return fileRelevance(pkg.Name, fi) < fileRelevance(pkg.Name, fj)
	})

	var docs []string
	for _, f := range filesWithDocs {
		docs = append(docs, f.Doc.Text())
	}
	return docs
}

// fileRelevance returns a sort key: lower is more relevant.
func fileRelevance(pkgName, filename string) int {
	switch {
	case path.Base(filename) == "doc.go":
		return 0
	case strings.TrimSuffix(path.Base(filename), path.Ext(filename)) == pkgName:
		return 1
	case !strings.HasSuffix(filename, "_test.go"):
		return 2
	default:
		return 3
	}
}
