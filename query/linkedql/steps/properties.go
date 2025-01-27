package steps

import (
	"fmt"

	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Properties{})
}

var _ linkedql.PathStep = (*Properties)(nil)

// Properties corresponds to .properties().
type Properties struct {
	From  linkedql.PathStep      `json:"from"`
	Names *linkedql.PropertyPath `json:"names"`
}

// Description implements Step.
func (s *Properties) Description() string {
	return "adds tags for all properties of the current entity"
}

func resolveNames(names *linkedql.PropertyPath) (linkedql.PropertyIRIs, error) {
	if names == nil {
		return nil, fmt.Errorf("Not implemented: should tag all properties")
	}
	switch n := names.PropertyPathI.(type) {
	case linkedql.PropertyStep:
		return nil, fmt.Errorf("Not implemented: should use step to resolve to properties")
	case linkedql.PropertyIRIs:
		return n, nil
	case linkedql.PropertyIRIStrings:
		return n.PropertyIRIs(), nil
	case linkedql.PropertyIRI:
		return linkedql.PropertyIRIs{n}, nil
	case linkedql.PropertyIRIString:
		return linkedql.PropertyIRIs{linkedql.PropertyIRI(n)}, nil
	default:
		return nil, fmt.Errorf("Unexpected type")
	}
}

// BuildPath implements linkedql.PathStep.
func (s *Properties) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	p := fromPath
	names, err := resolveNames(s.Names)
	if err != nil {
		return nil, err
	}
	for _, n := range names {
		name := quad.IRI(n).FullWith(ns)
		tag := string(name)
		p = p.Save(name, tag)
	}
	return p, nil
}
