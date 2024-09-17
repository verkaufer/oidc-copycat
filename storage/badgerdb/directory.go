package badgerdb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dgraph-io/badger/v4"
	oidc_copycat "github.com/verkaufer/oidc-copycat"
)

const (
	USER_KEY_PREFIX = "user:"
)

func (db *DB) GetUser(userId string) (*oidc_copycat.User, error) {

	txn := db.db.NewTransaction(false)
	defer txn.Discard()

	key := []byte(USER_KEY_PREFIX + userId)

	item, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	_, err = item.ValueCopy(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to read user data: %w", err)
	}

	// TODO: serialize to oidc_copycat.User

	return nil, nil
}

func (db *DB) GetAllUsers() ([]oidc_copycat.User, error) {
	// https://github.com/alexedwards/scs/blob/master/badgerstore/badgerstore.go#L92
	return nil, nil
}

func (db *DB) CreateUser(u *oidc_copycat.User) (*oidc_copycat.User, error) {

	serializedUser, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("failed to Marshal: %w", err)
	}

	txn := db.db.NewTransaction(true)
	defer txn.Discard()

	// TODO: Key needs to be hashed, lowercase email which will also be the "fallback" identifier
	key := []byte(USER_KEY_PREFIX + strings.ToLower(u.Email))

	entry := badger.NewEntry(key, serializedUser)
	if err := txn.SetEntry(entry); err != nil {
		return nil, err
	}

	if err := txn.Commit(); err != nil {
		return nil, err
	}

	return u, nil
}

func (db *DB) UpdateUser(u *oidc_copycat.User) (*oidc_copycat.User, error) {
	return nil, nil
}

func (db *DB) DeleteUser(userId string) error {
	txn := db.db.NewTransaction(true)
	defer txn.Discard()

	key := []byte(USER_KEY_PREFIX + userId)

	err := txn.Delete(key)
	if err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}
