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
	linkedql.Register(&GreaterThan{})
}

var _ linkedql.PathStep = (*GreaterThan)(nil)

// GreaterThan corresponds to gt().
type GreaterThan struct {
	From  linkedql.PathStep `json:"from"`
	Value quad.Value        `json:"value"`
}

// Description implements Step.
func (s *GreaterThan) Description() string {
	return "Greater than equals filters out values that are not greater than given value"
}

// BuildPath implements linkedql.PathStep.
func (s *GreaterThan) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Filter(iterator.CompareGT, linkedql.AbsoluteValue(s.Value, ns)), nil
}
