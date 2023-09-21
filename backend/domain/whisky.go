package domain

import (
	"context"
	"net/http"
)

type Whisky struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ABV          int    `json:"abv"`
	Link         string `json:"link"`
	Description  string `json:"description"`
	Title        string `json:"title"`
	DistilleryId int    `json:"distilleryId"`
}

type WhiskyUseCase interface {
	Fetch(ctx context.Context) ([]*Whisky, error)
	GetByID(ctx context.Context, id int64) (*Whisky, error)
	// Update(ctx context.Context, whisky *Whisky) error
	// Store(context.Context, *Whisky) error
	// Delete(ctx context.Context, id int64) error
}

type Storer interface {
	Update(ctx context.Context, whisky *Whisky) error
	Store(context.Context, *Whisky) error
}

type Getter interface {
	Fetch(ctx context.Context) ([]*Whisky, error)
	GetByID(ctx context.Context, id int64) (*Whisky, error)
}

type Deleter interface {
	Delete(ctx context.Context, id int64) error
}

type WhiskyRepository interface {
	// Storer
	Getter
	// Deleter
}

type WhiskyController interface {
	GetAllWhiskies(w http.ResponseWriter, r *http.Request)
	GetWhiskyByID(w http.ResponseWriter, r *http.Request)
}

func NewWhisky(id int64, name string, abv int, link string, description string, title string) (*Whisky, error) {
	w := &Whisky{
		ID:          id,
		Name:        name,
		ABV:         abv,
		Link:        link,
		Description: description,
		Title:       title,
	}

	return w, nil
}
