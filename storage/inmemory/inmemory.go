package inmemory

type DB struct {
	store map[string][]byte
}

func New() *DB {
	return &DB{
		store: make(map[string][]byte),
	}
}
