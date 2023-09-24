package usecase

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
)

type commentUseCase struct {
	repo domain.CommentRepository
}

func New(repo *domain.CommentRepository) domain.CommentUseCase {
	return &commentUseCase{
		repo: *repo,
	}
}

func (uc *commentUseCase) GetByID(ctx context.Context, id int64) (*domain.Comment, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *commentUseCase) Store(ctx context.Context, comment *domain.Comment) error {
	return uc.repo.Store(ctx, comment)
}

func (uc *commentUseCase) GetByWhiskyID(ctx context.Context, whiskyId int64) ([]*domain.Comment, error) {
	return uc.repo.GetByWhiskyID(ctx, whiskyId)
}
