package visitors

import (
	"strings"

	"github.com/sourcegraph/scip-go/internal/funk"
	"github.com/sourcegraph/scip-go/internal/symbols"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"golang.org/x/tools/go/packages"
)

type Scope struct {
	descriptors []*scip.Descriptor
}

func NewScope(pkgPath string) *Scope {
	return &Scope{
		descriptors: []*scip.Descriptor{
			{
				Name:   pkgPath,
				Suffix: scip.Descriptor_Namespace,
			},
		},
	}
}

func (s *Scope) push(name string, suffix scip.Descriptor_Suffix) {
	s.descriptors = append(s.descriptors, &scip.Descriptor{Name: name, Suffix: suffix})
}

func (s *Scope) pop() {
	s.descriptors = s.descriptors[:len(s.descriptors)-1]
}

func (s *Scope) makeSymbol(pkg *packages.Package, name string, suffix scip.Descriptor_Suffix) string {
	return symbols.FromDescriptors(pkg, append(s.descriptors, &scip.Descriptor{Name: name, Suffix: suffix})...)
}

func (s *Scope) String() string {
	return strings.Join(funk.Map(s.descriptors, func(d *scip.Descriptor) string {
		return d.Name
	}), " > ")
}
