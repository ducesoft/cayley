package linkedql

import (
	"context"

	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/graph/iterator"
	"github.com/ducesoft/cayley/graph/refs"
	"github.com/ducesoft/cayley/quad"
	"github.com/ducesoft/cayley/quad/jsonld"
	"github.com/ducesoft/cayley/quad/voc"
	"github.com/ducesoft/cayley/query"
	"github.com/ducesoft/cayley/query/path"
)

var _ query.Iterator = (*ValueIterator)(nil)

// ValueIterator is an iterator of values from the graph.
type ValueIterator struct {
	Namer   refs.Namer
	path    *path.Path
	scanner iterator.Scanner
	err     error
}

// NewValueIterator returns a new ValueIterator for a path and namer.
func NewValueIterator(p *path.Path, namer refs.Namer) *ValueIterator {
	return &ValueIterator{Namer: namer, path: p}
}

// NewValueIteratorFromPathStep attempts to build a path from PathStep and return a new ValueIterator of it.
// If BuildPath fails returns error.
func NewValueIteratorFromPathStep(step PathStep, qs graph.QuadStore, ns *voc.Namespaces) (*ValueIterator, error) {
	p, err := step.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return NewValueIterator(p, qs), nil
}

// Next implements query.Iterator.
func (it *ValueIterator) Next(ctx context.Context) bool {
	if it.scanner == nil {
		it.scanner = it.path.BuildIterator(ctx).Iterate()
	}
	return it.scanner.Next(ctx)
}

// Value returns the current value
func (it *ValueIterator) Value() quad.Value {
	if it.scanner == nil {
		return nil
	}
	rname, err := it.Namer.NameOf(it.scanner.Result())
	if err != nil {
		it.err = err
	}
	return rname
}

// Result implements query.Iterator.
func (it *ValueIterator) Result() interface{} {
	// FIXME(iddan): only convert when collation is JSON/JSON-LD, leave as Ref otherwise
	return jsonld.FromValue(it.Value())
}

// Err implements query.Iterator.
func (it *ValueIterator) Err() error {
	if it.err != nil {
		return it.err
	}
	if it.scanner == nil {
		return nil
	}
	return it.scanner.Err()
}

// Close implements query.Iterator.
func (it *ValueIterator) Close() error {
	if it.scanner == nil {
		return nil
	}
	return it.scanner.Close()
}
