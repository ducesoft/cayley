package all

import (
	// import all implementations that hidalgo supports
	_ "github.com/ducesoft/cayley/dal/kv/all"

	// make sure to import kv package, so it can re-register hidalgo's backends
	_ "github.com/ducesoft/cayley/graph/kv"
	// legacy: override bolt implementation; check the package for details
	_ "github.com/ducesoft/cayley/graph/kv/bolt"
)
