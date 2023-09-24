package usecase

import (
	"context"
	"errors"

	"github.com/AndersStigsson/whisky-calendar/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo domain.UserRepository
}

func New(repo *domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		repo: *repo,
	}
}

func (uc *userUseCase) Login(ctx context.Context, user *domain.User) error {
	dbUser, err := uc.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return errors.New("Incorrect password or username")
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return errors.New("Incorrect password or username")
	}
	return nil
}

func (uc *userUseCase) Register(ctx context.Context, user *domain.User) error {
	dbUser, err := uc.repo.GetByUsername(ctx, user.Username)
	if dbUser != nil || err == nil {
		return errors.New("Username is already taken")
	}

	return uc.repo.Store(ctx, user)
}
