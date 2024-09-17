package badgerdb

import "github.com/dgraph-io/badger/v4"

type DB struct {
	db *badger.DB
}

func New(db *badger.DB) *DB {
	return &DB{
		db: db,
	}
}
