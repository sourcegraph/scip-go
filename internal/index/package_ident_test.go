package index

import (
	"go/ast"
	"go/token"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindPackageDocs(t *testing.T) {
	type FileInfo struct {
		Name    string
		DocText string
	}

	makePackage := func(pkgName string, fileInfo []FileInfo) *packages.Package {
		fset := token.NewFileSet()
		var syntax []*ast.File

		for _, info := range fileInfo {
			var doc *ast.CommentGroup
			if info.DocText != "" {
				doc = &ast.CommentGroup{
					List: []*ast.Comment{{Text: "// " + info.DocText}},
				}
			}

			tf := fset.AddFile(info.Name, fset.Base(), 1)

			syntax = append(syntax, &ast.File{
				Doc:     doc,
				Package: tf.Pos(0),
				Name:    &ast.Ident{Name: pkgName},
			})
		}

		return &packages.Package{
			Name:    pkgName,
			PkgPath: "test-package",
			Fset:    fset,
			Syntax:  syntax,
		}
	}

	t.Run("returns nil when no file has docs", func(t *testing.T) {
		pkg := makePackage("smol", []FileInfo{
			{"smol.go", ""},
			{"other.go", ""},
		})
		if docs := findPackageDocs(pkg); docs != nil {
			t.Errorf("expected nil, got %v", docs)
		}
	})

	t.Run("returns single doc", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"mylib.go", ""},
			{"has_docs.go", "Package docs"},
		})
		docs := findPackageDocs(pkg)
		if len(docs) != 1 {
			t.Fatalf("expected 1 doc, got %d", len(docs))
		}
	})

	t.Run("returns all docs sorted with doc.go first", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"mylib.go", "from mylib"},
			{"doc.go", "from doc.go"},
			{"other.go", "from other"},
		})
		docs := findPackageDocs(pkg)
		if len(docs) != 3 {
			t.Fatalf("expected 3 docs, got %d", len(docs))
		}
		if !strings.Contains(docs[0], "from doc.go") {
			t.Errorf("expected doc.go first, got %q", docs[0])
		}
		if !strings.Contains(docs[1], "from mylib") {
			t.Errorf("expected mylib.go second, got %q", docs[1])
		}
	})

	t.Run("returns all docs sorted with package name match first when no doc.go", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"other.go", "from other"},
			{"mylib.go", "from mylib"},
		})
		docs := findPackageDocs(pkg)
		if len(docs) != 2 {
			t.Fatalf("expected 2 docs, got %d", len(docs))
		}
		if !strings.Contains(docs[0], "from mylib") {
			t.Errorf("expected mylib.go first, got %q", docs[0])
		}
	})

	t.Run("returns nil for empty syntax", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{})
		if docs := findPackageDocs(pkg); docs != nil {
			t.Errorf("expected nil, got %v", docs)
		}
	})
}
