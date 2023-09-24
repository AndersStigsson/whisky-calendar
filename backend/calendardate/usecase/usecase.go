package usecase

import (
	"context"

	"github.com/AndersStigsson/whisky-calendar/domain"
)

type calendarDateUseCase struct {
	repo domain.CalendarDateRepository
}

func NewCalendarDateUseCase(repo *domain.CalendarDateRepository) domain.CalendarDateUseCase {
	return &calendarDateUseCase{
		repo: *repo,
	}
}

func (uc *calendarDateUseCase) Fetch(ctx context.Context) ([]*domain.CalendarDate, error) {
	return uc.repo.Fetch(ctx)
}

func (uc *calendarDateUseCase) GetByDayOfMonth(ctx context.Context, doy int64) (*domain.CalendarDate, error) {
	return uc.repo.GetByDayOfMonth(ctx, doy)
}
