package leveldb

import (
	"testing"

	"github.com/ducesoft/cayley/dal/kv/flat"
	"github.com/ducesoft/cayley/dal/kv/kvtest"
)

func TestLeveldb(t *testing.T) {
	kvtest.RunTestLocal(t, flat.UpgradeOpenPath(OpenPath), nil)
}
