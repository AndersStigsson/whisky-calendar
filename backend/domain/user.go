package domain

import (
	"context"
	"net/http"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Name     string `json:"name"`
}

func NewUser(id int64, username string, password string, name string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
		Name:     name,
	}
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	// Login(ctx context.Context, username string, password string) error
	Store(ctx context.Context, user *User) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}

type UserUseCase interface {
	Login(ctx context.Context, user *User) (*User, error)
	Register(ctx context.Context, user *User) (*User, error)
}

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}
