package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Order{})
}

var _ linkedql.PathStep = (*Order)(nil)

// Order corresponds to .order().
type Order struct {
	From linkedql.PathStep `json:"from"`
}

// Description implements Step.
func (s *Order) Description() string {
	return "sorts the results in ascending order according to the current entity / value"
}

// BuildPath implements linkedql.PathStep.
func (s *Order) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Order(), nil
}
