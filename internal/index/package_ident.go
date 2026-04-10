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
		return fileRelevance(pkg, filesWithDocs[i]) < fileRelevance(pkg, filesWithDocs[j])
	})

	var docs []string
	for _, f := range filesWithDocs {
		docs = append(docs, f.Doc.Text())
	}
	return docs
}

// fileRelevance returns a sort key: lower is more relevant.
func fileRelevance(pkg *packages.Package, f *ast.File) int {
	fPath := pkg.Fset.Position(f.Pos()).Filename

	switch {
	case path.Base(fPath) == "doc.go":
		return 0
	case strings.TrimSuffix(path.Base(fPath), path.Ext(fPath)) == pkg.Name:
		return 1
	case !strings.HasSuffix(fPath, "_test.go"):
		return 2
	default:
		return 3
	}
}
