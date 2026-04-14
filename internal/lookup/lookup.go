package lookup

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"log/slog"
	"sync"

	"github.com/scip-code/scip/bindings/go/scip"
	"github.com/sourcegraph/scip-go/internal/newtypes"
	"github.com/sourcegraph/scip-go/internal/symbols"
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
		symbols:    map[newtypes.PackageID]*Package{},
		pkgSymbols: map[newtypes.PackageID]string{},
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
	if original, ok := p.fields[pos]; ok {
		if original.Symbol == symbol.Symbol {
			return
		}
		slog.Warn(fmt.Sprintf(
			"[scip.lookup] Conflicting symbol at %v: %s vs %s",
			p.pkg.Fset.Position(pos), original.Symbol, symbol.Symbol))
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

type Global struct {
	m          sync.Mutex
	symbols    map[newtypes.PackageID]*Package
	pkgSymbols map[newtypes.PackageID]string
}

func (p *Global) Add(pkgSymbols *Package) {
	p.m.Lock()
	p.symbols[newtypes.GetID(pkgSymbols.pkg)] = pkgSymbols
	p.m.Unlock()
}

func (p *Global) SetPkgSymbol(pkg *packages.Package) string {
	sym := symbols.FromDescriptors(pkg, &scip.Descriptor{
		Name:   pkg.PkgPath,
		Suffix: scip.Descriptor_Namespace,
	})
	p.m.Lock()
	p.pkgSymbols[newtypes.GetID(pkg)] = sym
	p.m.Unlock()
	return sym
}

func (p *Global) GetPkgSymbolByID(pkgID newtypes.PackageID) (string, bool) {
	sym, ok := p.pkgSymbols[pkgID]
	return sym, ok
}

func (p *Global) GetPkgSymbol(pkg *packages.Package) (string, bool) {
	return p.GetPkgSymbolByID(newtypes.GetID(pkg))
}

func (p *Global) GetPackage(pkg *packages.Package) *Package {
	return p.symbols[newtypes.GetID(pkg)]
}

var (
	skippedTypes   = map[string]struct{}{}
	builtinSymbols = map[string]*scip.SymbolInformation{}
)

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

	return nil, false, errors.New(fmt.Sprintf(
		"failed to create symbol for obj: %T %+v\n%s",
		obj,
		obj,
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
