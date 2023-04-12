package all

import (
	_ "github.com/ducesoft/cayley/dal/kv/flat/badger"
	_ "github.com/ducesoft/cayley/dal/kv/flat/btree"
	_ "github.com/ducesoft/cayley/dal/kv/flat/leveldb"
	_ "github.com/ducesoft/cayley/dal/kv/flat/pebble"
)
