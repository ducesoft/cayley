package mysql_test

import (
	"testing"

	_ "github.com/ducesoft/cayley/dal/tuple/sql/mysql/test"

	"github.com/ducesoft/cayley/dal/tuple/sql/mysql"
	"github.com/ducesoft/cayley/dal/tuple/sql/sqltest"
)

func TestMySQL(t *testing.T) {
	sqltest.Test(t, mysql.Name)
}

func BenchmarkMySQL(b *testing.B) {
	sqltest.Benchmark(b, mysql.Name)
}
