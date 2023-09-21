package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type whiskyRepository struct {
	db *sqlx.DB
}

func NewMySQLWhiskyRepository(db *sqlx.DB) domain.WhiskyRepository {
	return &whiskyRepository{
		db: db,
	}
}

type WhiskyModel struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	ABV          int    `db:"abv"`
	Link         string `db:"link"`
	Description  string `db:"description"`
	Title        string `db:"title"`
	DistilleryId int    `db:"distillery_id"`
}

// func (r *whiskyRepository) Update(ctx context.Context, whisky *domain.Whisky) error {
// }

// func (r *whiskyRepository) Store(context.Context, *domain.Whisky) error {
// }

func (r *whiskyRepository) Fetch(ctx context.Context) ([]*domain.Whisky, error) {
	var ww []*WhiskyModel
	rows, err := r.db.QueryxContext(ctx, getAllWhiskies)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "whiskyRepo.Fetch.QueryxContext")
		}
		return []*domain.Whisky{}, nil
	}

	defer rows.Close()
	for rows.Next() {
		w := &WhiskyModel{}
		if err = rows.StructScan(w); err != nil {
			return nil, errors.Wrap(err, "whiskyRepo.Fetch.Structscan")
		}
		ww = append(ww, w)
	}

	var dww []*domain.Whisky
	for _, w := range ww {
		dw, err := w.TranslateToDomain()
		if err != nil {
			fmt.Printf("Do something with this: %v", err)
		}
		dww = append(dww, dw)
	}

	return dww, nil
}

func (r *whiskyRepository) GetByID(ctx context.Context, id int64) (*domain.Whisky, error) {
	w := &WhiskyModel{}
	if err := r.db.GetContext(ctx, w, getWhiskyByID, id); err != nil {
		return nil, errors.Wrap(err, "whiskyRepo.GetByID.GetContext")
	}

	return w.TranslateToDomain()
}

// func (r *whiskyRepository) Delete(ctx context.Context, id int64) error {
// 	return errors.Wrap(errors.New("Not yet implemented"), "whiskyRepo.Delete")
// }

func (w *WhiskyModel) TranslateToDomain() (*domain.Whisky, error) {
	return domain.NewWhisky(
		w.ID,
		w.Name,
		w.ABV,
		w.Link,
		w.Description,
		w.Title,
	)
}
