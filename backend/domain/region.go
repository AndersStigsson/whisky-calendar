package domain

import (
	"context"
	"net/http"
)

type Region struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewRegion(id int64, name string) *Region {
	return &Region{
		ID:   id,
		Name: name,
	}
}

type RegionUseCase interface {
	GetByID(ctx context.Context, id int64) (*Region, error)
}

type RegionRepository interface {
	GetByID(ctx context.Context, id int64) (*Region, error)
}

type RegionController interface {
	GetRegionByID(w http.ResponseWriter, r *http.Request)
}
