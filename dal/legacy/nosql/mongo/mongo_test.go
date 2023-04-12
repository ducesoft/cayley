package mongo_test

import (
	"testing"

	"github.com/ducesoft/cayley/dal/legacy/nosql/mongo"
	_ "github.com/ducesoft/cayley/dal/legacy/nosql/mongo/test"
	"github.com/ducesoft/cayley/dal/legacy/nosql/nosqltest"
)

func TestMongo(t *testing.T) {
	nosqltest.Test(t, mongo.Name)
}

func BenchmarkMongo(b *testing.B) {
	nosqltest.Benchmark(b, mongo.Name)
}
