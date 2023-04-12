package postgres_test

import (
	"testing"

	_ "github.com/ducesoft/cayley/dal/tuple/sql/postgres/test"

	"github.com/ducesoft/cayley/dal/tuple/sql/postgres"
	"github.com/ducesoft/cayley/dal/tuple/sql/sqltest"
)

func TestPostgreSQL(t *testing.T) {
	sqltest.Test(t, postgres.Name)
}

func BenchmarkPostgreSQL(b *testing.B) {
	sqltest.Benchmark(b, postgres.Name)
}
