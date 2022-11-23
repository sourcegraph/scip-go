package lookup

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"sync"

	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

func NewPackageSymbols(pkg *packages.Package) *Package {
	return &Package{
		pkg:    pkg,
		fields: map[token.Pos]*scip.SymbolInformation{},
	}
}

func NewGlobalSymbols() *Global {
	return &Global{
		symbols: map[string]*Package{},
	}
}

type Package struct {
	pkg    *packages.Package
	fields map[token.Pos]*scip.SymbolInformation
}

func (p *Package) Set(pos token.Pos, symbol *scip.SymbolInformation) {
	// TODO: Could remove this once we are 100% confident we're not overlapping...
	if original, ok := p.fields[pos]; ok {
		if original != symbol {
			panic(fmt.Sprintf("Cannot add pos to new symbol: %s %s", original, symbol))
		}
	}

	p.fields[pos] = symbol
}

func (p *Package) Get(pos token.Pos) (*scip.SymbolInformation, bool) {
	field, ok := p.fields[pos]
	return field, ok
}

func (p *Package) GetSymbol(pos token.Pos) (string, bool) {
	field, ok := p.Get(pos)
	if ok && field != nil {
		return field.Symbol, true
	} else {
		return "", false
	}
}

// TODO: Don't love that this copies everything... :'(
func (p *Package) Symbols() []*scip.SymbolInformation {
	symbols := make([]*scip.SymbolInformation, 0, len(p.fields))
	for _, symbol := range p.fields {
		symbols = append(symbols, symbol)
	}
	return symbols
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
	case *types.Const:
		return "", false, nil
	}

	symbol, ok = symbols.FromObject(pkg, obj)
	if !ok {
		panic(fmt.Sprintf("failed to create symbol for builtin obj: %s %T %+v", pkg.Name, obj, obj))
	}

	return symbol, true, nil
}

func (p *Global) GetSymbolInformation(pkg *packages.Package, pos token.Pos) (*scip.SymbolInformation, bool) {
	pkgFields, ok := p.symbols[pkg.PkgPath]
	if !ok {
		fmt.Println("whoa whoa whoa... missing package?", pkg.PkgPath)
		return nil, false
	}

	field, ok := pkgFields.Get(pos)
	return field, ok
}

func (p *Global) GetSymbol(pkg *packages.Package, pos token.Pos) (string, bool) {
	field, ok := p.GetSymbolInformation(pkg, pos)
	if ok && field != nil {
		return field.Symbol, true
	} else {
		return "", false
	}
}
