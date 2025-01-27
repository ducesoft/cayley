// Copyright 2017 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leveldb

import (
	"os"

	hkv "github.com/ducesoft/cayley/dal/kv"
	"github.com/ducesoft/cayley/dal/kv/flat"
	"github.com/ducesoft/cayley/dal/kv/flat/leveldb"
	"github.com/ducesoft/cayley/graph"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const (
	Type = leveldb.Name
)

func Create(path string, m graph.Options) (hkv.KV, error) {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		return nil, err
	}
	db, err := leveldb.Open(path, &opt.Options{
		ErrorIfExist: true,
	})
	if os.IsExist(err) {
		return nil, graph.ErrDatabaseExists
	} else if err != nil {
		return nil, err
	}
	nosync, _ := m.BoolKey("nosync", false)
	db.SetWriteOptions(&opt.WriteOptions{
		Sync: !nosync,
	})
	return flat.Upgrade(db), nil
}

func Open(path string, m graph.Options) (hkv.KV, error) {
	db, err := leveldb.Open(path, &opt.Options{
		ErrorIfMissing: true,
	})
	if err != nil {
		return nil, err
	}
	nosync, _ := m.BoolKey("nosync", false)
	db.SetWriteOptions(&opt.WriteOptions{
		Sync: !nosync,
	})
	return flat.Upgrade(db), nil
}
