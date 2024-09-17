package inmemory

import (
	"encoding/json"
	"fmt"

	oidc_copycat "github.com/verkaufer/oidc-copycat"
)

const (
	userKeyPrefix = "user:"
)

func (db *DB) GetUser(userId string) (*oidc_copycat.User, error) {

	key := fmt.Sprintf("%s%s", userKeyPrefix, userId)

	storedUser, found := db.store[key]
	if !found {
		return nil, fmt.Errorf("user not found")
	}

	u := &oidc_copycat.User{}
	if err := json.Unmarshal(storedUser, u); err != nil {
		return nil, fmt.Errorf("failed to Unmarshal: %w", err)
	}

	return u, nil
}

func (db *DB) ListUsers() ([]oidc_copycat.User, error) {
	return nil, nil
}

func (db *DB) CreateUser(u *oidc_copycat.User) (*oidc_copycat.User, error) {
	return nil, nil
}

func (db *DB) DeleteUser(userId string) error {
	return nil
}
