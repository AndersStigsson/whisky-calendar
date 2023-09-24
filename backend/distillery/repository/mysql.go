package repository

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type distilleryRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.DistilleryRepository {
	return &distilleryRepository{
		db: db,
	}
}

type DistilleryModel struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	RegionID int    `db:"region_id"`
}

func (r *distilleryRepository) GetByID(ctx context.Context, id int64) (*domain.Distillery, error) {
	dm := &DistilleryModel{}
	if err := r.db.GetContext(ctx, dm, getDistilleryByID, id); err != nil {
		return nil, errors.Wrap(err, "distilleryRepo.GetByID.GetContext")
	}

	return dm.TranslateToDomain()
}

func (dm *DistilleryModel) TranslateToDomain() (*domain.Distillery, error) {
	dd := domain.NewDistillery(
		dm.ID,
		dm.Name,
		dm.RegionID,
	)

	return dd, nil
}
