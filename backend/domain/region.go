package domain

import "context"

type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RegionUseCase interface {
	GetByID(ctx context.Context, id int64) (*Region, error)
}

type RegionRepository interface {
	GetByID(ctx context.Context, id int64) (*Region, error)
}
