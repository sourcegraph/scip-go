package index

import (
	"go/ast"
	"go/token"
	"slices"
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

	// doc produces the text that ast.CommentGroup.Text() returns for a "// text" comment.
	doc := func(text string) string {
		return text + "\n"
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

	t.Run("returns nil for empty syntax", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{})
		if docs := findPackageDocs(pkg); docs != nil {
			t.Errorf("expected nil, got %v", docs)
		}
	})

	t.Run("returns single doc", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"mylib.go", ""},
			{"has_docs.go", "Package docs"},
		})
		want := []string{doc("Package docs")}
		got := findPackageDocs(pkg)
		if !slices.Equal(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("returns all docs sorted with doc.go first", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"mylib.go", "from mylib"},
			{"doc.go", "from doc.go"},
			{"other.go", "from other"},
		})
		want := []string{doc("from doc.go"), doc("from mylib"), doc("from other")}
		got := findPackageDocs(pkg)
		if !slices.Equal(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("returns all docs sorted with package name match first when no doc.go", func(t *testing.T) {
		pkg := makePackage("mylib", []FileInfo{
			{"other.go", "from other"},
			{"mylib.go", "from mylib"},
		})
		want := []string{doc("from mylib"), doc("from other")}
		got := findPackageDocs(pkg)
		if !slices.Equal(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func TestFileRelevance(t *testing.T) {
	tests := []struct {
		filename string
		want     int
	}{
		{"doc.go", 0},
		{"mylib.go", 1},
		{"other.go", 2},
		{"other_test.go", 3},
	}
	for _, tt := range tests {
		if got := fileRelevance("mylib", tt.filename); got != tt.want {
			t.Errorf("fileRelevance(%q) = %d, want %d", tt.filename, got, tt.want)
		}
	}
}
