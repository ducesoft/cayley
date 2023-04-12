package elastic_test

import (
	"testing"

	"github.com/ducesoft/cayley/dal/legacy/nosql/elastic"
	_ "github.com/ducesoft/cayley/dal/legacy/nosql/elastic/test"
	"github.com/ducesoft/cayley/dal/legacy/nosql/nosqltest"
)

func TestElastic(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	nosqltest.Test(t, elastic.Name)
}

func BenchmarkElastic(b *testing.B) {
	nosqltest.Benchmark(b, elastic.Name)
}
