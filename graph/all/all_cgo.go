//go:build cgo
// +build cgo

package all

import (
	// backends requiring cgo
	_ "github.com/ducesoft/cayley/graph/sql/sqlite"
)
