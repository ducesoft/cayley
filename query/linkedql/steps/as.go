package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&As{})
}

var _ linkedql.PathStep = (*As)(nil)

// As corresponds to .tag().
type As struct {
	From linkedql.PathStep `json:"from"`
	Name string            `json:"name"`
}

// Description implements Step.
func (s *As) Description() string {
	return "assigns the resolved values of the from step to a given name. The name can be used with the Select and Documents steps to retrieve the values or to return to the values in further steps with the Back step. It resolves to the values of the from step."
}

// BuildPath implements linkedql.PathStep.
func (s *As) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Tag(s.Name), nil
}
