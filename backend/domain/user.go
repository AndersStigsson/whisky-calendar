package domain

import "context"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*User, error)
	Login(ctx context.Context, username string, password string) error
	Register(ctx context.Context, username string, password string) error
}

type UserUseCase interface {
	Login(ctx context.Context, username string, password string) error
	Register(ctx context.Context, username string, password string) error
}
