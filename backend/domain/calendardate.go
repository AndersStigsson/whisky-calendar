package domain

import (
	"context"
	"net/http"
	"time"
)

type CalendarDate struct {
	DayOfMonth int        `json:"dayOfMonth"`
	RevealDate time.Time  `json:"revealDate"`
	WhiskyID   int        `json:"whiskyId"`
	Comments   *[]Comment `json:"comments"`
}

type CalendarDateUseCase interface {
	Fetch(ctx context.Context) ([]*CalendarDate, error)
	GetByDayOfMonth(ctx context.Context, doy int64) (*CalendarDate, error)
}

type CalendarDateGetter interface {
	Fetch(ctx context.Context) ([]*CalendarDate, error)
	GetByDayOfMonth(ctx context.Context, doy int64) (*CalendarDate, error)
}

type CalendarDateRepository interface {
	CalendarDateGetter
}

type CalendarDateController interface {
	GetAllDates(w http.ResponseWriter, r *http.Request)
	GetDateByDayOfMonth(w http.ResponseWriter, r *http.Request)
}

func NewCalendarDate(dayOfMonth int, revealDate time.Time, whiskyID int, comments *[]Comment) (*CalendarDate, error) {
	w := &CalendarDate{
		DayOfMonth: dayOfMonth,
		RevealDate: revealDate,
		WhiskyID:   whiskyID,
		Comments:   comments,
	}

	return w, nil
}
