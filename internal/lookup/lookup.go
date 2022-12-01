package lookup

import (
	"errors"
	"fmt"
	"go/ast"
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
		symbols:  map[string]*Package{},
		pkgNames: map[string]*PackageName{},
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

type PackageName struct {
	Symbol *scip.SymbolInformation
	Pos    token.Pos
}

type Global struct {
	m        sync.Mutex
	symbols  map[string]*Package
	pkgNames map[string]*PackageName
}

func (p *Global) Add(pkgSymbols *Package) {
	p.m.Lock()
	p.symbols[pkgSymbols.pkg.PkgPath] = pkgSymbols
	p.m.Unlock()
}

func (p *Global) SetPkgName(pkg *packages.Package, pkgDeclaration *ast.File) {
	p.m.Lock()
	p.pkgNames[pkg.PkgPath] = &PackageName{
		Symbol: &scip.SymbolInformation{
			Symbol: symbols.FromDescriptors(pkg, &scip.Descriptor{
				Name:   pkg.PkgPath,
				Suffix: scip.Descriptor_Namespace,
			}),
			Documentation: []string{},
			Relationships: []*scip.Relationship{},
		},
		Pos: pkgDeclaration.Name.NamePos,
	}
	p.m.Unlock()
}

func (p *Global) GetPkgNameSymbol(pkgPath string) *PackageName {
	return p.pkgNames[pkgPath]
}

func (p *Global) GetPackage(pkg *packages.Package) *Package {
	return p.symbols[pkg.PkgPath]
}

var skippedTypes = map[string]struct{}{}
var builtinSymbols = map[string]*scip.SymbolInformation{}

// GetSymbolOfObject returns a symbol and whether we were successful at finding.
//
// We can return an empty string if this object should be ignored.
func (p *Global) GetSymbolOfObject(obj types.Object) (*scip.SymbolInformation, bool, error) {
	if _, ok := skippedTypes[obj.Id()]; ok {
		return nil, false, nil
	}

	if sym, ok := builtinSymbols[obj.Id()]; ok {
		return sym, true, nil
	}

	switch obj := obj.(type) {
	case *types.PkgName:
		panic(fmt.Sprintf("should never lookup PkgName %s | %+v", obj.Id(), obj.Imported().Path()))
	case *types.Nil:
		return nil, false, nil
	}

	pkg := obj.Pkg()
	if pkg == nil {
		switch obj := obj.(type) {
		case *types.TypeName:
			skippedTypes[obj.Id()] = struct{}{}
			return nil, false, nil
		case *types.Const:
			return nil, false, nil
		case *types.Builtin:
			return nil, false, nil
		}

		panic(fmt.Sprintf("failed to create symbol for builtin obj: %T %+v | %s", obj, obj, obj.Id()))
	}

	pkgPath := pkg.Path()
	symbol, ok := p.getSymbolInformationByPath(pkgPath, obj.Pos())
	if ok {
		return symbol, true, nil
	}

	switch obj := obj.(type) {
	case *types.Var:
		// , "| position", pkg.Fset.Position(obj.Pos())))
		return nil, false, errors.New(fmt.Sprintln("obj", obj, "| origin", obj.Origin()))
	}

	panic(fmt.Sprintf("failed to create symbol for obj: %T %+v", obj, obj))

}

func (p *Global) getSymbolInformationByPath(pkgPath string, pos token.Pos) (*scip.SymbolInformation, bool) {
	pkgFields, ok := p.symbols[pkgPath]
	if !ok {
		fmt.Println("whoa whoa whoa... missing package?", pkgPath)
		return nil, false
	}

	field, ok := pkgFields.Get(pos)
	return field, ok
}

func (p *Global) GetSymbolInformation(pkg *packages.Package, pos token.Pos) (*scip.SymbolInformation, bool) {
	return p.getSymbolInformationByPath(pkg.PkgPath, pos)
}

func (p *Global) GetSymbol(pkg *packages.Package, pos token.Pos) (string, bool) {
	field, ok := p.GetSymbolInformation(pkg, pos)
	if ok && field != nil {
		return field.Symbol, true
	} else {
		return "", false
	}
}
