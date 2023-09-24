package repository

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *sqlx.DB
}

func NewMySQL(db *sqlx.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

type UserModel struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	um := &UserModel{}
	if err := r.db.GetContext(ctx, um, getUserByID, id); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetByID.GetContext")
	}

	return um.TranslateToDomain()
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	um := &UserModel{}
	if err := r.db.GetContext(ctx, um, getUserByUsername, username); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetByID.GetContext")
	}

	return um.TranslateToDomain()
}

func (r *userRepository) Store(ctx context.Context, user *domain.User) error {
	u := &UserModel{}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err := r.db.QueryRowxContext(
		ctx,
		createUser,
		user.Username,
		password,
		user.Name,
	).StructScan(u); err != nil {
		return err
	}

	return nil
}

func (um *UserModel) TranslateToDomain() (*domain.User, error) {
	du := domain.NewUser(
		um.ID,
		um.Username,
		um.Password,
		um.Name,
	)

	return du, nil
}
