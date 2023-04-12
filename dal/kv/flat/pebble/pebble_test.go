package pebble

import (
	"testing"

	"github.com/ducesoft/cayley/dal/kv/flat"
	"github.com/ducesoft/cayley/dal/kv/kvtest"
)

func TestPebble(t *testing.T) {
	kvtest.RunTestLocal(t, flat.UpgradeOpenPath(OpenPath), &kvtest.Options{
		NoTx: true,
	})
}
