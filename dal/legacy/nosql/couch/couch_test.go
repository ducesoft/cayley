//go:build !js
// +build !js

package couch_test

import (
	"testing"

	"github.com/ducesoft/cayley/dal/legacy/nosql/couch"
	_ "github.com/ducesoft/cayley/dal/legacy/nosql/couch/test"
	"github.com/ducesoft/cayley/dal/legacy/nosql/nosqltest"
)

func TestCouch(t *testing.T) {
	nosqltest.Test(t, couch.NameCouch)
}

func BenchmarkCouch(b *testing.B) {
	nosqltest.Benchmark(b, couch.NameCouch)
}
