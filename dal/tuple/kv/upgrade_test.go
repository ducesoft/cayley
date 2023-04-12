package tuplekv_test

import (
	"testing"

	"github.com/ducesoft/cayley/dal/kv/flat"
	"github.com/ducesoft/cayley/dal/kv/flat/btree"
	"github.com/ducesoft/cayley/dal/tuple"
	tuplekv "github.com/ducesoft/cayley/dal/tuple/kv"
	"github.com/ducesoft/cayley/dal/tuple/tupletest"
)

func TestKV2Tuple(t *testing.T) {
	tupletest.RunTest(t, func(t testing.TB) tuple.Store {
		kdb := btree.New()
		db := tuplekv.New(flat.Upgrade(kdb))
		return db
	}, &tupletest.Options{
		NoLocks: true,
	})
}
