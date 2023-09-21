package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	"github.com/AndersStigsson/whisky-calendar/domain"
)

type whiskyController struct {
	service domain.WhiskyUseCase
	ctx     context.Context
	router  router.Router
}

func NewWhiskyController(service domain.WhiskyUseCase, ctx context.Context, router router.Router) domain.WhiskyController {
	return &whiskyController{
		service: service,
		ctx:     ctx,
		router:  router,
	}
}

func (wc *whiskyController) GetAllWhiskies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ww, err := wc.service.Fetch(wc.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("Couldn't fetch anything"))
	}
	fmt.Printf("%v\n", ww)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ww)
}

func (wc *whiskyController) GetWhiskyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := wc.router.ParsePathVariable(r, "id")
	ww, _ := wc.service.GetByID(wc.ctx, id.(int64))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ww)
}
