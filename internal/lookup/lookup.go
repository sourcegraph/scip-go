package lookup

import (
	"go/token"
	"sync"

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

func (p *Global) GetSymbol(pkg *packages.Package, pos token.Pos) (string, bool) {
	pkgFields, ok := p.symbols[pkg.PkgPath]
	if !ok {
		return "", false
	}

	field, ok := pkgFields.Get(pos)
	return field, ok
}
