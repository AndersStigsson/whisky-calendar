package repository

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type regionRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.RegionRepository {
	return &regionRepository{
		db: db,
	}
}

type RegionModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (r *regionRepository) GetByID(ctx context.Context, id int64) (*domain.Region, error) {
	rm := &RegionModel{}
	if err := r.db.GetContext(ctx, rm, getRegionByID, id); err != nil {
		return nil, errors.Wrap(err, "regionRepo.GetByID.GetContext")
	}

	return rm.TranslateToDomain()
}

func (rm *RegionModel) TranslateToDomain() (*domain.Region, error) {
	dr := domain.NewRegion(
		rm.ID,
		rm.Name,
	)

	return dr, nil
}
