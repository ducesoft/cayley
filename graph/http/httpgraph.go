package httpgraph

import (
	"net/http"

	"github.com/ducesoft/cayley/graph"
)

type QuadStore interface {
	graph.QuadStore
	ForRequest(r *http.Request) (graph.QuadStore, error)
}
