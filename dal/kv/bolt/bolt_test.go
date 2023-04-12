package bolt

import (
	"path/filepath"
	"testing"

	"github.com/ducesoft/cayley/dal/kv"
	"github.com/ducesoft/cayley/dal/kv/kvtest"
)

func TestBolt(t *testing.T) {
	kvtest.RunTestLocal(t, func(path string) (kv.KV, error) {
		path = filepath.Join(path, "bolt.db")
		return OpenPath(path)
	}, nil)
}
