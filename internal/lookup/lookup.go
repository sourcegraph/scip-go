package lookup

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"sync"

	"github.com/sourcegraph/scip-go/internal/symbols"
	"golang.org/x/tools/go/packages"
)

func NewPackageSymbols(pkg *packages.Package) *Package {
	return &Package{
		pkg:    pkg,
		fields: map[token.Pos]string{},
	}
}

func NewGlobalSymbols() *Global {
	return &Global{
		symbols: map[string]*Package{},
	}
}

type Package struct {
	pkg    *packages.Package
	fields map[token.Pos]string
}

func (p *Package) Set(pos token.Pos, symbol string) {
	// TODO: Could remove this once we are 100% confident we're not overlapping...
	if original, ok := p.fields[pos]; ok {
		if original != symbol {
			panic(fmt.Sprintf("Cannot add pos to new symbol: %s %s", original, symbol))
		}
	}

	p.fields[pos] = symbol
}

func (p *Package) Get(pos token.Pos) (string, bool) {
	field, ok := p.fields[pos]
	return field, ok
}

type Global struct {
	m       sync.Mutex
	symbols map[string]*Package
}

func (p *Global) Add(pkgSymbols *Package) {
	p.m.Lock()
	p.symbols[pkgSymbols.pkg.PkgPath] = pkgSymbols
	p.m.Unlock()
}

func (p *Global) GetPackage(pkg *packages.Package) *Package {
	return p.symbols[pkg.PkgPath]
}

var skippedTypes = map[string]struct{}{}

// GetSymbolOfObject returns a symbol and whether we were successful at finding.
//
// We can return an empty string if this object should be ignored.
func (p *Global) GetSymbolOfObject(pkg *packages.Package, obj types.Object) (string, bool, error) {
	switch obj.(type) {
	case *types.PkgName:
		// TODO: Should emit something for this I think
		return "", false, nil
	case *types.Nil:
		return "", false, nil
	}

	symbol, ok := p.GetSymbol(pkg, obj.Pos())
	if ok {
		return symbol, true, nil
	}

	switch obj := obj.(type) {
	case *types.Var:
		return "", false, errors.New(fmt.Sprintln("obj", obj, "| origion", obj.Origin(), "| position", pkg.Fset.Position(obj.Pos())))
	}

	if pkg.Name != "builtin" {
		panic(fmt.Sprintf("non-builtin package failing: %s | %T | %+v", pkg.Name, obj, obj))
	}

	switch obj := obj.(type) {
	case *types.TypeName:
		skippedTypes[obj.Name()] = struct{}{}

		return "", false, nil
	}

	symbol, ok = symbols.FromObject(pkg, obj)
	if !ok {
		panic(fmt.Sprintf("failed to create symbol for builtin obj: %s %T %+v", pkg.Name, obj, obj))
	}

	return symbol, true, nil
}

func (p *Global) GetSymbol(pkg *packages.Package, pos token.Pos) (string, bool) {
	pkgFields, ok := p.symbols[pkg.PkgPath]
	if !ok {
		fmt.Println("whoa whoa whoa... missing package?", pkg.PkgPath)
		return "", false
	}

	field, ok := pkgFields.Get(pos)
	return field, ok
}
