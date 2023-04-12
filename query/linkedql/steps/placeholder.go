package steps

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query/linkedql"
	"github.com/ducesoft/cayley/query/path"
)

func init() {
	linkedql.Register(&Placeholder{})
}

var _ linkedql.PathStep = (*Placeholder)(nil)

// Placeholder corresponds to .Placeholder().
type Placeholder struct{}

// Description implements Step.
func (s *Placeholder) Description() string {
	return "is like Vertex but resolves to the values in the context it is placed in. It should only be used where a linkedql.PathStep is expected and can't be resolved on its own."
}

// BuildPath implements linkedql.PathStep.
func (s *Placeholder) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	return path.StartMorphism(), nil
}
