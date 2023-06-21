package visitors

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"github.com/sourcegraph/scip-go/internal/document"
	"github.com/sourcegraph/scip-go/internal/handler"
	"github.com/sourcegraph/scip-go/internal/loader"
	"github.com/sourcegraph/scip-go/internal/lookup"
	"github.com/sourcegraph/scip-go/internal/newtypes"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

const symbolDefinition = int32(scip.SymbolRole_Definition)
const symbolReference = int32(scip.SymbolRole_ReadAccess)

func NewFileVisitor(
	doc *document.Document,
	pkg *packages.Package,
	file *ast.File,
	pkgLookup loader.PackageLookup,
	pkgSymbols *lookup.Package,
	globalSymbols *lookup.Global,
) *fileVisitor {
	caseClauses := map[token.Pos]types.Object{}
	for implicit, obj := range pkg.TypesInfo.Implicits {
		if _, ok := implicit.(*ast.CaseClause); ok {
			caseClauses[obj.Pos()] = obj
		}
	}

	// Package occurrence always goes into the list of occurrences for a document
	occurrences := []*scip.Occurrence{
		doc.PackageOccurrence,
	}

	return &fileVisitor{
		doc:           doc,
		pkg:           pkg,
		file:          file,
		pkgLookup:     pkgLookup,
		locals:        map[token.Pos]string{},
		pkgSymbols:    pkgSymbols,
		globalSymbols: globalSymbols,
		occurrences:   occurrences,
		overrides: struct {
			caseClauses     map[token.Pos]types.Object
			pkgNameOverride map[newtypes.PackageID]string
		}{
			caseClauses:     caseClauses,
			pkgNameOverride: map[newtypes.PackageID]string{},
		},
	}
}

// fileVisitor visits an entire file, but it must be called
// after StructVisitor.
//
// Iterates over a file,
type fileVisitor struct {
	// Document to append occurrences to
	doc *document.Document

	// Current file information
	pkg  *packages.Package
	file *ast.File

	// soething
	pkgLookup loader.PackageLookup

	// local definition position to symbol
	locals map[token.Pos]string

	// field definition position to symbol for the package
	pkgSymbols *lookup.Package

	// field definition position to symbol for the entire compliation
	globalSymbols *lookup.Global

	// occurrences in this file
	occurrences []*scip.Occurrence

	// Overriding Definition Behvaior:
	overrides struct {
		// Case clauses have to map particular positions to different types
		caseClauses map[token.Pos]types.Object

		// maps tokens for package declaration to a local var,
		// if ImportSpec.Name is not nil. Otherwise, just use package directly
		pkgNameOverride map[newtypes.PackageID]string
	}
}

// Implements ast.Visitor
var _ ast.Visitor = &fileVisitor{}

func (f *fileVisitor) createNewLocalSymbol(pos token.Pos) string {
	if _, ok := f.locals[pos]; ok {
		panic("Cannot create a new local symbol for an ident that has already been created")
	}

	f.locals[pos] = fmt.Sprintf("local %d", len(f.locals))
	return f.locals[pos]
}

func (v *fileVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.ImportSpec:
		// Generate import references
		importedPackage := v.pkg.Imports[strings.Trim(node.Path.Value, `"`)]
		if importedPackage == nil {
			fmt.Println("Could not find: ", node.Path)
			return nil
		}

		if node.Name != nil && node.Name.Name != "." {
			sym := v.createNewLocalSymbol(node.Name.Pos())
			v.NewDefinition(sym, symbols.RangeFromName(v.pkg.Fset.Position(node.Name.Pos()), node.Name.Name, false))

			// Save package name override, so that we use the new local symbol
			// within this file
			v.overrides.pkgNameOverride[newtypes.GetID(importedPackage)] = sym
		}

		position := v.pkg.Fset.Position(node.Path.Pos())
		v.emitImportReference(v.globalSymbols, v.doc, position, importedPackage)

		return nil

	case *ast.SelectorExpr:
		if ident, ok := node.X.(*ast.Ident); ok {
			use := v.pkg.TypesInfo.Uses[ident]

			// We special case handling PkgNames because they do some goofy things
			// compared to almost every other construct in the language.
			switch sel := use.(type) {
			case *types.PkgName:
				pos := ident.NamePos
				position := v.pkg.Fset.Position(pos)
				pkgID := newtypes.GetFromTypesPackage(sel.Imported())

				var symbolName string
				if overrideSymbol, ok := v.overrides.pkgNameOverride[pkgID]; ok {
					symbolName = overrideSymbol
				} else {
					symbol := v.globalSymbols.GetPkgNameSymbolByID(pkgID)
					if symbol == nil {
						handler.ErrOrPanic("Missing symbol for package: %s", sel.Imported().Path())
						return nil
					}

					symbolName = symbol.Symbol
				}

				v.AppendSymbolReference(symbolName, scipRange(position, sel), nil)

				// Then walk the selection
				ast.Walk(v, node.Sel)

				// and since we've handled the rest, end visit
				return nil
			}
		}

		return v
	case *ast.File:
		if node.Doc != nil {
			ast.Walk(v, node.Doc)
		}

		// Handle package name declaration separately
		// No need to: ast.Walk(v, n.Name)

		walkDeclList(v, node.Decls)
		return nil
	case *ast.Ident:
		// Short circuit if this is a blank identifier
		if node.Name == "_" {
			return nil
		}

		pos := node.NamePos
		position := v.pkg.Fset.Position(pos)

		// Short circuit on case clauses
		if obj, ok := v.overrides.caseClauses[node.Pos()]; ok {
			sym := v.createNewLocalSymbol(obj.Pos())
			v.NewDefinition(sym, scipRange(position, obj))
			return nil
		}

		info := v.pkg.TypesInfo

		// Emit Definition
		def := info.Defs[node]
		if def != nil {
			var sym string
			if pkgSymbols, ok := v.pkgSymbols.GetSymbol(def.Pos()); ok {
				sym = pkgSymbols
			} else if globalSymbol, ok := v.globalSymbols.GetSymbol(v.pkg, def.Pos()); ok {
				sym = globalSymbol
			} else {
				sym = v.createNewLocalSymbol(def.Pos())
			}

			v.NewDefinition(sym, scipRange(position, def))
		}

		// Emit Reference
		ref := info.Uses[node]
		if ref != nil {
			var (
				symbol       string
				overrideType types.Type
			)

			if localSymbol, ok := v.locals[ref.Pos()]; ok {
				symbol = localSymbol

				if _, ok := v.overrides.caseClauses[ref.Pos()]; ok {
					overrideType = v.pkg.TypesInfo.TypeOf(node)
				}
			} else {
				var err error
				symInfo, ok, err := v.globalSymbols.GetSymbolOfObject(ref)
				if err != nil {
					// _, ok := v.pkgLookup[symbols.PkgPathFromObject(ref)]
					// if !ok {
					// 	if err := handler.ErrOrPanic(
					// 		"Failed to find a package for ref: |%+v|\nNode: %s",
					// 		ref,
					// 		v.pkg.Fset.Position(node.Pos()),
					// 	); err != nil {
					// 		return v
					// 	}
					//
					// }

					if err := handler.ErrOrPanic(
						"Unable to find symbol of object: %s\nNode Position -> %s\n\nPath: %s\n\n",
						err,
						v.pkg.Fset.Position(node.Pos()),
						ref.Pkg().Path(),
					); err != nil {
						return v
					}
				}

				if !ok {
					return v
				}

				// Set the resulting info
				symbol = symInfo.Symbol
			}

			v.AppendSymbolReference(symbol, scipRange(position, ref), overrideType)
		}

		if def == nil && ref == nil {
			handler.ErrOrPanic(
				"Neither def nor ref found: %s | %T | %s",
				node.Name,
				node,
				v.pkg.Fset.Position(node.Pos()),
			)
		}
	}

	return v
}

func (v *fileVisitor) emitImportReference(
	globalSymbols *lookup.Global,
	doc *document.Document,
	position token.Position,
	importedPackage *packages.Package,
) {
	scipRange := symbols.RangeFromName(position, importedPackage.PkgPath, true)
	symbol := globalSymbols.GetPkgNameSymbol(importedPackage)
	if symbol == nil {
		handler.ErrOrPanic("Missing symbol for package path: %s", importedPackage.ID)
		return
	}

	if symbol == nil {
		handler.ErrOrPanic("Missing symbol information for package: %s", importedPackage.ID)
		return
	}

	v.AppendSymbolReference(symbol.Symbol, scipRange, nil)
}

// NewDefinition emits a scip.Occurence ONLY. This will not emit a
// new symbol. You must do that using DeclareNewSymbol[ForPos]
func (v *fileVisitor) NewDefinition(symbol string, rng []int32) {
	v.occurrences = append(v.occurrences, &scip.Occurrence{
		Range:       rng,
		Symbol:      symbol,
		SymbolRoles: symbolDefinition,
	})
}

func (v *fileVisitor) AppendSymbolReference(symbol string, rng []int32, overrideType types.Type) {
	var documentation []string = nil
	if overrideType != nil {
		tyString := overrideType.String()
		if tyString != "" {
			documentation = append(documentation, symbols.FormatCode(tyString))
		}
	}

	v.occurrences = append(v.occurrences, &scip.Occurrence{
		Range:                 rng,
		Symbol:                symbol,
		SymbolRoles:           symbolReference,
		OverrideDocumentation: documentation,
	})
}

func (v *fileVisitor) ToScipDocument() *scip.Document {
	documentFile := v.pkg.Fset.File(v.file.Pos())
	if documentFile == nil {
		panic("that shouldn't happend")
	}

	documentSymbols := v.pkgSymbols.SymbolsForFile(documentFile)
	return &scip.Document{
		Language:     "go",
		RelativePath: v.doc.RelativePath,
		Occurrences:  v.occurrences,
		Symbols:      documentSymbols,
	}
}
