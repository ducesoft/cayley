package badger

import (
	"testing"

	"github.com/ducesoft/cayley/dal/kv/flat"
	"github.com/ducesoft/cayley/dal/kv/kvtest"
)

func TestBadger(t *testing.T) {
	kvtest.RunTestLocal(t, flat.UpgradeOpenPath(OpenPath), nil)
}
