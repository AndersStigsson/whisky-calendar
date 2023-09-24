package domain

import (
	"context"
	"net/http"
)

type Distillery struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	RegionID int    `json:"regionId"`
}

func NewDistillery(id int64, name string, regionID int) *Distillery {
	return &Distillery{
		ID:       id,
		Name:     name,
		RegionID: regionID,
	}
}

type DistilleryUseCase interface {
	GetByID(ctx context.Context, id int64) (*Distillery, error)
}

type DistilleryRepository interface {
	GetByID(ctx context.Context, id int64) (*Distillery, error)
}

type DistilleryController interface {
	GetDistilleryByID(w http.ResponseWriter, r *http.Request)
}
