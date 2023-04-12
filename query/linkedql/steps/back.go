package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Back{})
}

var _ linkedql.PathStep = (*Back)(nil)

// Back corresponds to .back().
type Back struct {
	From linkedql.PathStep `json:"from"`
	Name string            `json:"name"`
}

// Description implements Step.
func (s *Back) Description() string {
	return "resolves to the values of the previous the step or the values assigned to name in a former step."
}

// BuildPath implements linkedql.PathStep.
func (s *Back) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Back(s.Name), nil
}
