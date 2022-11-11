package index

// func traverseFields(pkgs []*packages.Package) *GlobalSymbols {
// 	ch := make(chan func())
//
// 	projectFields := NewGlobalSymbols()
// 	go func() {
// 		defer close(ch)
//
// 		for _, pkg := range pkgs {
// 			// Bind pkg
// 			pkg := pkg
//
// 			ch <- func() {
// 				packageFields := NewPackageSymbols(pkg)
//
// 				visitor := StructVisitor{
// 					mod:    pkg.Module,
// 					Fields: packageFields,
// 					curScope: []*scip.Descriptor{
// 						{
// 							Name:   pkg.PkgPath,
// 							Suffix: scip.Descriptor_Namespace,
// 						},
// 					},
// 				}
//
// 				for _, f := range pkg.Syntax {
// 					ast.Walk(visitor, f)
// 				}
//
// 				projectFields.add(&packageFields)
// 			}
//
// 		}
// 	}()
//
// 	n := uint64(len(pkgs))
// 	wg, count := parallel.Run(ch)
// 	output.WithProgressParallel(
// 		wg,
// 		"Traversing Field Definitions",
// 		output.Options{
// 			Verbosity:      output.DefaultOutput,
// 			ShowAnimations: true,
// 		},
// 		count,
// 		n,
// 	)
//
// 	return &projectFields
// }
