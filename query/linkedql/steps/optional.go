package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Optional{})
}

var _ linkedql.PathStep = (*Optional)(nil)

// Optional corresponds to .optional().
type Optional struct {
	From linkedql.PathStep `json:"from"`
	Step linkedql.PathStep `json:"step"`
}

// Description implements Step.
func (s *Optional) Description() string {
	return "attempts to follow the given path from the current entity / value, if fails the entity / value will still be kept in the results"
}

// BuildPath implements linkedql.PathStep.
func (s *Optional) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	p, err := s.Step.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Optional(p), nil
}
