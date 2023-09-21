package usecase

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
)

type whiskyUseCase struct {
	repo domain.WhiskyRepository
}

func NewWhiskyUseCase(repo *domain.WhiskyRepository) domain.WhiskyUseCase {
	return &whiskyUseCase{
		repo: *repo,
	}
}

func (uc *whiskyUseCase) Fetch(ctx context.Context) ([]*domain.Whisky, error) {
	return uc.repo.Fetch(ctx)
}

func (uc *whiskyUseCase) GetByID(ctx context.Context, id int64) (*domain.Whisky, error) {
	return uc.repo.GetByID(ctx, id)
}
