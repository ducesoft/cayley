package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Unique{})
}

var _ linkedql.PathStep = (*Unique)(nil)

// Unique corresponds to .unique().
type Unique struct {
	From linkedql.PathStep `json:"from"`
}

// Description implements Step.
func (s *Unique) Description() string {
	return "removes duplicate values from the path."
}

// BuildPath implements linkedql.PathStep.
func (s *Unique) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Unique(), nil
}
