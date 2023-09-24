package usecase

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
)

type regionUseCase struct {
	repo domain.RegionRepository
}

func New(repo *domain.RegionRepository) domain.RegionUseCase {
	return &regionUseCase{
		repo: *repo,
	}
}

func (r *regionUseCase) GetByID(ctx context.Context, id int64) (*domain.Region, error) {
	return r.repo.GetByID(ctx, id)
}
