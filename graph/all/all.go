package all

import (
	// supported backends
	_ "github.com/ducesoft/cayley/graph/kv/all"
	_ "github.com/ducesoft/cayley/graph/memstore"
	_ "github.com/ducesoft/cayley/graph/nosql"
	_ "github.com/ducesoft/cayley/graph/sql/cockroach"
	_ "github.com/ducesoft/cayley/graph/sql/mysql"
	_ "github.com/ducesoft/cayley/graph/sql/postgres"
)
