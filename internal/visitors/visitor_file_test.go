package visitors

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"golang.org/x/tools/go/packages"
)

func TestToScipDocument_IncludesLocalSymbols(t *testing.T) {
	src := `
	package main

	func main() {
		x := 42
		y := "hello"
		if true {
			z := x + 1
			_ = z
		}
		_ = y
	}
	`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatalf("Failed to parse source: %v", err)
	}

	// Create a mock package
	pkg := &packages.Package{
		Fset:   fset,
		Syntax: []*ast.File{file},
		Types:  types.NewPackage("main", "main"),
		TypesInfo: &types.Info{
			Implicits: make(map[ast.Node]types.Object),
		},
	}

	doc := &document.Document{RelativePath: "test.go"}

	pkgSymbols := lookup.NewPackageSymbols(pkg)
	globalSymbols := lookup.NewGlobalSymbols()
	pkgLookup := make(loader.PackageLookup) // empty

	visitor := NewFileVisitor(
		doc, pkg, file, pkgLookup, pkgSymbols, globalSymbols)

	// Simulate visiting the file to create local symbols
	// In real usage, ast.Walk would populate the locals
	// For testing, we'll manually add some local symbols to simulate what happens
	visitor.locals = map[token.Pos]string{
		token.Pos(100): "local 0", // x
		token.Pos(200): "local 1", // y
		token.Pos(300): "local 2", // z
	}

	scipDoc := visitor.ToScipDocument()
	localSymbolCount := 0
	for _, symbol := range scipDoc.Symbols {
		if len(symbol.Symbol) >= 5 && symbol.Symbol[:5] == "local" {
			localSymbolCount++
		}
	}

	if localSymbolCount != 3 {
		t.Errorf(
			"Expected 3 local symbols in Document.symbols, got %d", localSymbolCount)
	}

	expectedLocals := map[string]bool{
		"local 0": false,
		"local 1": false,
		"local 2": false,
	}

	for _, symbol := range scipDoc.Symbols {
		if _, ok := expectedLocals[symbol.Symbol]; ok {
			expectedLocals[symbol.Symbol] = true
		}
	}

	for localSymbol, found := range expectedLocals {
		if !found {
			t.Errorf(
				"Expected local symbol %q not found in Document.symbols", localSymbol)
		}
	}
}
