package bitcask

import (
	"encoding/json"
	"fmt"
	"strings"

	oidc_copycat "github.com/verkaufer/oidc-copycat"
	"go.mills.io/bitcask/v2"
)

const (
	USER_KEY_PREFIX = "user:"
)

func (db *DB) GetUser(userId string) (*oidc_copycat.User, error) {

	key := bitcask.Key(USER_KEY_PREFIX + userId)

	item, err := db.store.Get(key)
	if err != nil {
		return nil, err
	}

	user := &oidc_copycat.User{}
	if err := json.Unmarshal(item, user); err != nil {
		return nil, fmt.Errorf("failed to Unmarshal: %w", err)
	}
	return user, nil
}

func (db *DB) GetAllUsers() ([]oidc_copycat.User, error) {

	return nil, nil
}

func (db *DB) CreateUser(u *oidc_copycat.User) (*oidc_copycat.User, error) {

	serializedUser, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("failed to Marshal: %w", err)
	}

	// TODO: Key needs to be hashed, lowercase email which will also be the "fallback" identifier
	key := []byte(USER_KEY_PREFIX + strings.ToLower(u.Email))

	if err := db.store.Put(key, serializedUser); err != nil {
		return nil, fmt.Errorf("failed to Put: %w", err)
	}

	return u, nil
}

func (db *DB) UpdateUser(u *oidc_copycat.User) (*oidc_copycat.User, error) {
	return nil, nil
}

func (db *DB) DeleteUser(userId string) error {

	key := []byte(USER_KEY_PREFIX + userId)

	err := db.store.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
