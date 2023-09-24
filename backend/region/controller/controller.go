package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	"github.com/AndersStigsson/whisky-calendar/domain"
)

type regionController struct {
	service domain.RegionUseCase
	ctx     context.Context
	router  router.Router
}

func New(service domain.RegionUseCase, ctx context.Context, router router.Router) domain.RegionController {
	return &regionController{
		service: service,
		ctx:     ctx,
		router:  router,
	}
}

func (c *regionController) GetRegionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := c.router.ParsePathVariable(r, "id")
	idInt, _ := strconv.ParseInt(id.(string), 10, 64)
	region, _ := c.service.GetByID(c.ctx, idInt)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(region)
}
