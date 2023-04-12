package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/graph/iterator"
	"github.com/ducesoft/cayley/quad"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&LessThanEquals{})
}

var _ linkedql.PathStep = (*LessThanEquals)(nil)

// LessThanEquals corresponds to lte().
type LessThanEquals struct {
	From  linkedql.PathStep `json:"from"`
	Value quad.Value        `json:"value"`
}

// Description implements Step.
func (s *LessThanEquals) Description() string {
	return "Less than equals filters out values that are not less than or equal given value"
}

// BuildPath implements linkedql.PathStep.
func (s *LessThanEquals) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Filter(iterator.CompareLTE, linkedql.AbsoluteValue(s.Value, ns)), nil
}
