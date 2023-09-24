package usecase

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
)

type distilleryUseCase struct {
	repo domain.DistilleryRepository
}

func New(repo *domain.DistilleryRepository) domain.DistilleryUseCase {
	return &distilleryUseCase{
		repo: *repo,
	}
}

func (d *distilleryUseCase) GetByID(ctx context.Context, id int64) (*domain.Distillery, error) {
	return d.repo.GetByID(ctx, id)
}
