//go:build docker
// +build docker

package nosql

import (
	"testing"

	hnosqltest "github.com/ducesoft/cayley/dal/legacy/nosql/nosqltest"
	_ "github.com/ducesoft/cayley/dal/legacy/nosql/nosqltest/all"
)

func TestNoSQL(t *testing.T) {
	hnosqltest.RunTest(t, nosqltest.TestAll)
}

func BenchmarkNoSQL(t *testing.B) {
	hnosqltest.RunBenchmark(t, nosqltest.BenchmarkAll)
}
