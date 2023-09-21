package domain

import "context"

type Distillery struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	RegionID int    `json:"regionId"`
}

type DistilleryUseCase interface {
	GetByID(ctx context.Context, id int64) (*Distillery, error)
}

type DistilleryRepository interface {
	GetByID(ctx context.Context, id int64) (*Distillery, error)
}
