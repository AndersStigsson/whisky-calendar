package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	"github.com/AndersStigsson/whisky-calendar/domain"
)

type distilleryController struct {
	service domain.DistilleryUseCase
	ctx     context.Context
	router  router.Router
}

func New(service domain.DistilleryUseCase, ctx context.Context, router router.Router) domain.DistilleryController {
	return &distilleryController{
		service: service,
		ctx:     ctx,
		router:  router,
	}
}

func (d *distilleryController) GetDistilleryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := d.router.ParsePathVariable(r, "id")
	idInt, _ := strconv.ParseInt(id.(string), 10, 64)
	distillery, _ := d.service.GetByID(d.ctx, idInt)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(distillery)
}
