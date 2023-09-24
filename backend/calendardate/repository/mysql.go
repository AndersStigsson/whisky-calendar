package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type calendarDateRepository struct {
	db *sqlx.DB
}

func NewMySQLCalendarDateRepository(db *sqlx.DB) domain.CalendarDateRepository {
	return &calendarDateRepository{
		db: db,
	}
}

type CalendarDateModel struct {
	DayOfMonth int       `db:"day_of_month"`
	RevealDate time.Time `db:"reveal_date"`
	WhiskyID   int       `db:"whisky_id"`
	Comments   *[]domain.Comment
}

func (r *calendarDateRepository) Fetch(ctx context.Context) ([]*domain.CalendarDate, error) {
	var cdcd []*CalendarDateModel
	rows, err := r.db.QueryxContext(ctx, getAllCalendarDates)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "calendarDateRepo.Fetch.QueryxContext")
		}
		return []*domain.CalendarDate{}, nil
	}

	defer rows.Close()
	for rows.Next() {
		cd := &CalendarDateModel{}
		if err = rows.StructScan(cd); err != nil {
			return nil, errors.Wrap(err, "calendarDateRepo.Fetch.Structscan")
		}
		cdcd = append(cdcd, cd)
	}

	var dcdcd []*domain.CalendarDate
	for _, cd := range cdcd {
		dcd, err := cd.TranslateToDomain()
		if err != nil {
			fmt.Printf("Do something with this: %v", err)
		}
		dcdcd = append(dcdcd, dcd)
	}

	return dcdcd, nil
}

func (r *calendarDateRepository) GetByDayOfMonth(ctx context.Context, doy int64) (*domain.CalendarDate, error) {
	cd := &CalendarDateModel{}
	if err := r.db.GetContext(ctx, cd, getDateByDayOfMonth, doy); err != nil {
		return nil, errors.Wrap(err, "calendarDateRepo.GetByDayOfMonth.GetContext")
	}

	return cd.TranslateToDomain()
}

func (cd *CalendarDateModel) TranslateToDomain() (*domain.CalendarDate, error) {
	return domain.NewCalendarDate(
		cd.DayOfMonth,
		cd.RevealDate,
		cd.WhiskyID,
		cd.Comments,
	)
}
