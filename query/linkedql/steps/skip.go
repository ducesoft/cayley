package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Skip{})
}

var _ linkedql.PathStep = (*Skip)(nil)

// Skip corresponds to .skip().
type Skip struct {
	From   linkedql.PathStep `json:"from"`
	Offset int64             `json:"offset"`
}

// Description implements Step.
func (s *Skip) Description() string {
	return "skips a number of nodes for current path."
}

// BuildPath implements linkedql.PathStep.
func (s *Skip) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Skip(s.Offset), nil
}
