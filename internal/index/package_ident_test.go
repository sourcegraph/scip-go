package index

import (
	"go/ast"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindPackageDocFile(t *testing.T) {
	type FileInfo struct {
		Name string
		Docs bool
	}

	makePackage := func(pkgName string, fileInfo []FileInfo) (*packages.Package, map[*ast.File]string) {
		fset := token.NewFileSet()

		syntax := []*ast.File{}
		posMap := map[*ast.File]string{}

		for _, info := range fileInfo {
			var doc *ast.CommentGroup
			if info.Docs {
				doc = &ast.CommentGroup{}
			}

			tf := fset.AddFile(info.Name, fset.Base(), 1)

			f := &ast.File{
				Doc:     doc,
				Package: tf.Pos(0),
				Name: &ast.Ident{
					Name: pkgName,
					Obj:  &ast.Object{},
				},
			}

			syntax = append(syntax, f)
			posMap[f] = info.Name
		}

		return &packages.Package{
			ID:      "test-package",
			Name:    pkgName,
			PkgPath: "test-package",
			Imports: map[string]*packages.Package{},
			Fset:    fset,
			Syntax:  syntax,
		}, posMap
	}

	t.Run("returns nil when no file has docs", func(t *testing.T) {
		pkg, _ := makePackage("smol", []FileInfo{
			{"smol.go", false},
			{"other.go", false},
		})
		if f := findPackageDocFile(pkg); f != nil {
			t.Errorf("expected nil, got a file")
		}
	})

	t.Run("returns the single file with docs", func(t *testing.T) {
		pkg, names := makePackage("mylib", []FileInfo{
			{"mylib.go", false},
			{"has_docs.go", true},
		})
		f := findPackageDocFile(pkg)
		if f == nil {
			t.Fatal("expected a file, got nil")
		}
		if name := names[f]; name != "has_docs.go" {
			t.Errorf("want=has_docs.go have=%s", name)
		}
	})

	t.Run("prefers doc.go when multiple files have docs", func(t *testing.T) {
		pkg, names := makePackage("mylib", []FileInfo{
			{"mylib.go", true},
			{"doc.go", true},
		})
		f := findPackageDocFile(pkg)
		if f == nil {
			t.Fatal("expected a file, got nil")
		}
		if name := names[f]; name != "doc.go" {
			t.Errorf("want=doc.go have=%s", name)
		}
	})

	t.Run("prefers exact package name match when multiple files have docs", func(t *testing.T) {
		pkg, names := makePackage("mylib", []FileInfo{
			{"other.go", true},
			{"mylib.go", true},
		})
		f := findPackageDocFile(pkg)
		if f == nil {
			t.Fatal("expected a file, got nil")
		}
		if name := names[f]; name != "mylib.go" {
			t.Errorf("want=mylib.go have=%s", name)
		}
	})

	t.Run("returns nil for empty syntax", func(t *testing.T) {
		pkg, _ := makePackage("mylib", []FileInfo{})
		if f := findPackageDocFile(pkg); f != nil {
			t.Errorf("expected nil, got a file")
		}
	})
}
