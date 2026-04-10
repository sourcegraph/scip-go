package index

import (
	"go/ast"
	"go/token"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindPackageDoc(t *testing.T) {
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

	t.Run("returns empty when no file has docs", func(t *testing.T) {
		pkg := makePackage("smol", []FileInfo{
			{"smol.go", ""},
			{"other.go", ""},
		})
		if doc := findPackageDoc(pkg); doc != "" {
			t.Errorf("expected empty, got %q", doc)
		}
	})

	t.Run("returns doc from the single file with docs", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"mylib.go", ""},
			{"has_docs.go", "Package docs here"},
		})
		doc := findPackageDoc(pkg)
		if doc == "" {
			t.Fatal("expected doc text, got empty")
		}
	})

	t.Run("prefers doc.go when multiple files have docs", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"mylib.go", "from mylib"},
			{"doc.go", "from doc.go"},
		})
		doc := findPackageDoc(pkg)
		if doc == "" {
			t.Fatal("expected doc text, got empty")
		}
		if !strings.Contains(doc, "from doc.go") {
			t.Errorf("expected doc from doc.go, got %q", doc)
		}
	})

	t.Run("prefers exact package name match when multiple files have docs", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"other.go", "from other"},
			{"mylib.go", "from mylib"},
		})
		doc := findPackageDoc(pkg)
		if doc == "" {
			t.Fatal("expected doc text, got empty")
		}
		if !strings.Contains(doc, "from mylib") {
			t.Errorf("expected doc from mylib.go, got %q", doc)
		}
	})

	t.Run("returns empty for empty syntax", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{})
		if doc := findPackageDoc(pkg); doc != "" {
			t.Errorf("expected empty, got %q", doc)
		}
	})
}
