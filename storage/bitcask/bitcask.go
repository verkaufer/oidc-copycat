package bitcask

import "go.mills.io/bitcask/v2"

type DB struct {
	store *bitcask.Bitcask
}

func New(b *bitcask.Bitcask) *DB {
	return &DB{
		store: b,
	}
}
