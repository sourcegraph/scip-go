package index

import (
	"go/ast"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestIndexer_findBestPackageDefinitionPath(t *testing.T) {
	type FileInfo struct {
		Name string
		Docs bool
	}

	makePackage := func(pkgName string, fileInfo []FileInfo) (*packages.Package, map[*ast.File]string) {
		fset := token.NewFileSet()

		files := map[string]*ast.File{}
		syntax := []*ast.File{}
		posMap := map[*ast.File]string{}

		for idx, info := range fileInfo {
			var doc *ast.CommentGroup
			if info.Docs {
				doc = &ast.CommentGroup{}
			}

			f := &ast.File{
				Doc:     doc,
				Package: 0,
				Name: &ast.Ident{
					NamePos: token.Pos(idx),
					Name:    pkgName,
					Obj:     &ast.Object{},
				},
			}

			files[info.Name] = f
			syntax = append(syntax, f)
			posMap[f] = info.Name

			fset.AddFile(info.Name, fset.Base(), 1)
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

	makeTest := func(name, pkgName, expected string, fileInfo []FileInfo) {
		t.Run(name, func(t *testing.T) {
			pkg, names := makePackage(pkgName, fileInfo)

			pkgToken, _ := findBestPackageDefinitionPath(pkg)
			if name := names[pkgToken]; name != expected {
				t.Errorf("incorrect hover text documentation. want=%s have=%s", name, expected)
			}
		})
	}

	makeTest("Should find exact name match",
		"smol",
		"smol.go",
		[]FileInfo{
			{"smol.go", false},
			{"other.go", false},
		},
	)

	makeTest("Should return something even if nothing matches",
		"smol",
		"random.go",
		[]FileInfo{
			{"random.go", false},
		},
	)

	makeTest("Should not pick _test files if package is not a test package",
		"unreleated",
		"smol.go",
		[]FileInfo{
			{"smol.go", false},
			{"smol_test.go", false},
		},
	)

	makeTest("Pick whatever has documentation",
		"mylib",
		"has_docs.go",
		[]FileInfo{
			{"mylib.go", false},
			{"has_docs.go", true},
		},
	)

	makeTest("should pick a name that is a closer edit distance than one far away",
		"http_router",
		"httprouter.go",
		[]FileInfo{
			{"httprouter.go", false},
			{"httpother.go", false},
		},
	)

	makeTest("should prefer test packages over other packages if the package name has test suffix",
		"mylib_test",
		"mylib_test.go",
		[]FileInfo{
			{"mylib_test.go", false},
			{"mylib.go", false},
		},
	)
}
