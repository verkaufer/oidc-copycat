package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type DirectoryReader interface {
	GetUser(userId string) (*User, error)
	GetAllUsers() (*UsersDirectory, error)
}

type DirectoryWriter interface {
	CreateUser(u *User) (*User, error)
	// TODO: Update, Delete
}

type DirectoryReaderWriter interface {
	DirectoryReader
	DirectoryWriter
}

type User struct {
	Identifier string `json:"identifier"`
	FirstName  string `json:"givenName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
}

type UsersDirectory struct {
	Users []User `json:"users"`
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type LocalJsonRepository struct {
	filename string

	// TODO: add lock?
}

func (r *LocalJsonRepository) GetAllUsers() (*UsersDirectory, error) {

	file, err := os.Open(r.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	rawBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var directory UsersDirectory
	json.Unmarshal(rawBytes, &directory)

	return &directory, nil
}

func (r *LocalJsonRepository) GetUser(userId string) (*User, error) {
	directory, err := r.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to read user directory: %w", err)
	}

	for _, user := range directory.Users {
		if user.Identifier == userId {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with userID %s not found", userId)
}

func NewLocalJsonRepository(filename string) *LocalJsonRepository {
	// assert filename not empty
	if filename == "" {
		panic("filename is not defined")
	}

	_, err := os.Stat(filename)
	if errors.Is(err, fs.ErrNotExist) {
		panic(fmt.Sprintf("file %s does not exist", filename))
	}

	return &LocalJsonRepository{
		filename: filename,
	}
}
