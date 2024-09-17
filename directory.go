package oidc_copycat

import (
	"fmt"
)

type DirectoryReader interface {
	GetUser(userId string) (*User, error)
	GetAllUsers() ([]User, error)
}

type DirectoryWriter interface {
	CreateUser(u *User) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(userId string) error
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

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type DirectoryService struct {
	repo DirectoryReaderWriter
}

func NewDirectoryService(datasource DirectoryReaderWriter) *DirectoryService {
	return &DirectoryService{
		repo: datasource,
	}
}

func (d *DirectoryService) GetUser(userId string) (*User, error) {
	return d.repo.GetUser(userId)
}

func (d *DirectoryService) ListUsers() ([]User, error) {
	return d.repo.GetAllUsers()
}

func (d *DirectoryService) CreateUser(user *User) (*User, error) {
	// TODO: check if user exists first
	return d.repo.CreateUser(user)
}
