package visitors

import "github.com/sourcegraph/scip/bindings/go/scip"

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
