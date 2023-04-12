//go:build docker
// +build docker

package postgres

import (
	"testing"

	"github.com/ducesoft/cayley/graph"
	"github.com/ducesoft/cayley/graph/sql/sqltest"
	"github.com/ducesoft/cayley/internal/dock"
	"github.com/jackc/pgx/v5/pgconn"
)

func makePostgres(t testing.TB) (string, graph.Options, func()) {
	var conf dock.Config

	conf.Image = "postgres:9.5"
	conf.OpenStdin = true
	conf.Tty = true
	conf.Env = []string{`POSTGRES_PASSWORD=postgres`}

	addr, closer := dock.RunAndWait(t, conf, "5432", func(addr string) bool {
		conn, err := pgconn.Open(`postgres://postgres:postgres@` + addr + `/postgres?sslmode=disable`)
		if err != nil {
			return false
		}
		conn.Close()
		return true
	})
	addr = `postgres://postgres:postgres@` + addr + `/postgres?sslmode=disable`
	return addr, nil, func() {
		closer()
	}
}

var conf = &sqltest.Config{
	TimeRound: true,
	TimeInMcs: true,
}

func TestPostgres(t *testing.T) {
	sqltest.TestAll(t, Type, makePostgres, conf)
}

func BenchmarkPostgres(t *testing.B) {
	sqltest.BenchmarkAll(t, Type, makePostgres, conf)
}
