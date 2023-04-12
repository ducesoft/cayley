package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Limit{})
}

var _ linkedql.PathStep = (*Limit)(nil)

// Limit corresponds to .limit().
type Limit struct {
	From  linkedql.PathStep `json:"from"`
	Limit int64             `json:"limit"`
}

// Description implements Step.
func (s *Limit) Description() string {
	return "limits a number of nodes for current path."
}

// BuildPath implements linkedql.PathStep.
func (s *Limit) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Limit(s.Limit), nil
}
