package lookup

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"sync"

	"github.com/sourcegraph/scip-go/internal/newtypes"
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
		symbols:  map[newtypes.PackageID]*Package{},
		pkgNames: map[newtypes.PackageID]*PackageName{},
	}
}

type Package struct {
	pkg    *packages.Package
	fields map[token.Pos]*scip.SymbolInformation
}

func (p *Package) SymbolsForFile(file *token.File) []*scip.SymbolInformation {
	var documentSymbols []*scip.SymbolInformation = nil
	for pos, symbol := range p.fields {
		if p.pkg.Fset.File(pos) == file {
			documentSymbols = append(documentSymbols, symbol)
		}
	}

	return documentSymbols
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

type PackageName struct {
	Symbol *scip.SymbolInformation
	Pos    token.Pos
}

type Global struct {
	m        sync.Mutex
	symbols  map[newtypes.PackageID]*Package
	pkgNames map[newtypes.PackageID]*PackageName
}

func (p *Global) Add(pkgSymbols *Package) {
	p.m.Lock()
	p.symbols[newtypes.GetID(pkgSymbols.pkg)] = pkgSymbols
	p.m.Unlock()
}

func (p *Global) SetPkgName(pkg *packages.Package, pkgDeclaration *ast.File) {
	p.m.Lock()
	p.pkgNames[newtypes.GetID(pkg)] = &PackageName{
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

func (p *Global) GetPkgNameSymbolByID(pkgID newtypes.PackageID) *scip.SymbolInformation {
	named, ok := p.pkgNames[pkgID]
	if !ok {
		return nil
	}

	return named.Symbol
}

func (p *Global) GetPkgNameSymbol(pkg *packages.Package) *scip.SymbolInformation {
	return p.GetPkgNameSymbolByID(newtypes.GetID(pkg))
}

func (p *Global) GetPackage(pkg *packages.Package) *Package {
	return p.symbols[newtypes.GetID(pkg)]
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
		case *types.Func:
			if orig := obj.Origin(); orig != nil {
				name := orig.FullName()
				switch name {
				case "(error).Error":
					return nil, false, nil
				}
			}
		}

		panic(fmt.Sprintf("failed to create symbol for builtin obj: %T %+v | %s", obj, obj, obj.Id()))
	}

	pkgPath := pkg.Path()
	for _, combination := range testPackageCombinations(pkgPath) {
		symbol, _, ok := p.getSymbolInformationByPath(combination, obj.Pos())
		if ok {
			return symbol, true, nil
		}
	}

	switch obj := obj.(type) {
	case *types.Var:
		// , "| position", pkg.Fset.Position(obj.Pos())))
		return nil, false, errors.New(fmt.Sprintln("obj", obj, "| origin", obj.Origin()))
	}

	// if !foundPkg {
	// 	panic(fmt.Sprintf(
	// 		"missing package (%s) for symbol: %T %+v\n%s",
	// 		pkgPath,
	// 		obj,
	// 		obj,
	// 		pkgFields.pkg.Fset.Position(obj.Pos()),
	// 	))
	// }

	return nil, false, errors.New(fmt.Sprintf(
		"failed to create symbol for obj: %T %+v\n%s",
		obj,
		obj,
		// pkgFields.pkg.Fset.Position(obj.Pos()),
		pkgPath,
	))
}

func (p *Global) getSymbolInformationByPath(pkgID newtypes.PackageID, pos token.Pos) (*scip.SymbolInformation, *Package, bool) {
	pkgFields, ok := p.symbols[pkgID]
	if !ok {
		return nil, nil, false
	}

	field, ok := pkgFields.Get(pos)
	return field, pkgFields, ok
}

func (p *Global) GetSymbolInformation(pkg *packages.Package, pos token.Pos) (*scip.SymbolInformation, bool) {
	info, _, ok := p.getSymbolInformationByPath(newtypes.GetID(pkg), pos)
	return info, ok
}

func (p *Global) GetSymbol(pkg *packages.Package, pos token.Pos) (string, bool) {
	field, ok := p.GetSymbolInformation(pkg, pos)
	if ok && field != nil {
		return field.Symbol, true
	} else {
		return "", false
	}
}

func testPackageCombinations(pkgPath string) []newtypes.PackageID {
	return []newtypes.PackageID{
		newtypes.PackageID(pkgPath),
		newtypes.PackageID(fmt.Sprintf("%s.test", pkgPath)),
		newtypes.PackageID(fmt.Sprintf("%s [%s.test]", pkgPath, pkgPath)),
	}
}
