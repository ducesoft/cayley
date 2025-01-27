package cayley

import (
	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/graph/iterator"
	_ "github.com/ducesoft/cayley/graph/memstore"
	"github.com/ducesoft/cayley/quad"
	"github.com/ducesoft/cayley/query/path"
	_ "github.com/ducesoft/cayley/writer"
)

var (
	StartMorphism = path.StartMorphism
	StartPath     = path.StartPath

	NewTransaction = graph.NewTransaction
)

type Iterator = iterator.Shape
type QuadStore = graph.QuadStore
type QuadWriter = graph.QuadWriter

type Path = path.Path

type Handle struct {
	graph.QuadStore
	graph.QuadWriter
}

func (h *Handle) Close() error {
	err := h.QuadWriter.Close()
	h.QuadStore.Close()
	return err
}

func Triple(subject, predicate, object interface{}) quad.Quad {
	return Quad(subject, predicate, object, nil)
}

func Quad(subject, predicate, object, label interface{}) quad.Quad {
	return quad.Make(subject, predicate, object, label)
}

func NewGraph(name, dbpath string, opts graph.Options) (*Handle, error) {
	qs, err := graph.NewQuadStore(name, dbpath, opts)
	if err != nil {
		return nil, err
	}
	qw, err := graph.NewQuadWriter("single", qs, nil)
	if err != nil {
		return nil, err
	}
	return &Handle{qs, qw}, nil
}

func NewMemoryGraph() (*Handle, error) {
	return NewGraph("memstore", "", nil)
}
