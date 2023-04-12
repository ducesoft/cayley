//go:build js
// +build js

package couch

import (
	"github.com/ducesoft/cayley/dal/base"
	"github.com/ducesoft/cayley/dal/legacy/nosql"
	_ "github.com/go-kivik/pouchdb/v4" // The PouchDB driver
)

const (
	NamePouch   = "pouch"
	DriverPouch = "couch"
)

func init() {
	nosql.Register(nosql.Registration{
		Registration: base.Registration{
			Name: NamePouch, Title: DriverPouch,
			Local: true, Volatile: false,
		},
		Traits: Traits(),
		New:    CreatePouch, Open: OpenPouch,
	})
}

func CreatePouch(addr string, ns string, opt nosql.Options) (nosql.Database, error) {
	return Dial(true, DriverPouch, addr, ns, opt)
}

func OpenPouch(addr string, ns string, opt nosql.Options) (nosql.Database, error) {
	return Dial(false, DriverPouch, addr, ns, opt)
}
