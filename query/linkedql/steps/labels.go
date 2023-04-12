package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Labels{})
}

var _ linkedql.PathStep = (*Labels)(nil)

// Labels corresponds to .labels().
type Labels struct {
	From linkedql.PathStep `json:"from"`
}

// Description implements Step.
func (s *Labels) Description() string {
	return "gets the list of inbound and outbound quad labels"
}

// BuildPath implements linkedql.PathStep.
func (s *Labels) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Labels(), nil
}
