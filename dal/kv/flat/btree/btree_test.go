package btree

import (
	"testing"

	"github.com/ducesoft/cayley/dal/kv"
	"github.com/ducesoft/cayley/dal/kv/flat"
	"github.com/ducesoft/cayley/dal/kv/kvtest"
)

func TestBtree(t *testing.T) {
	kvtest.RunTest(t, func(t testing.TB) kv.KV {
		return flat.Upgrade(New())
	}, nil)
}
