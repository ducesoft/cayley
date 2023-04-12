package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Count{})
}

var _ linkedql.PathStep = (*Count)(nil)

// Count corresponds to .count().
type Count struct {
	From linkedql.PathStep `json:"from"`
}

// Description implements Step.
func (s *Count) Description() string {
	return "resolves to the number of the resolved values of the from step"
}

// BuildPath implements linkedql.PathStep.
func (s *Count) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Count(), nil
}
