package mysqltest

import (
	"testing"

	dockertest "github.com/ory/dockertest/v3"

	sqltuple "github.com/ducesoft/cayley/dal/tuple/sql"
	"github.com/ducesoft/cayley/dal/tuple/sql/postgres"
	"github.com/ducesoft/cayley/dal/tuple/sql/sqltest"
)

var versions = []string{
	"13",
}

func init() {
	var vers []sqltest.Version
	for _, v := range versions {
		vers = append(vers, sqltest.Version{
			Name: v, Factory: PostgresVersion(v),
		})
	}
	sqltest.Register(postgres.Name, vers...)
}

func PostgresVersion(vers string) sqltest.Database {
	const image = "postgres"
	return sqltest.Database{
		Recreate: false,
		Run: func(t testing.TB) string {
			pool, err := dockertest.NewPool("")
			if err != nil {
				t.Fatal(err)
			}

			cont, err := pool.Run(image, vers, []string{
				"POSTGRES_PASSWORD=postgres",
			})
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				_ = cont.Close()
			})

			const port = "5432/tcp"
			addr := `postgres://postgres:postgres@` + cont.GetHostPort(port)

			err = pool.Retry(func() error {
				cli, err := sqltuple.OpenSQL(postgres.Name, addr, "")
				if err != nil {
					return err
				}
				defer cli.Close()
				return cli.Ping()
			})
			if err != nil {
				t.Fatal(err)
			}
			return addr
		},
	}
}
