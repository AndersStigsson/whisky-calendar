package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	"github.com/AndersStigsson/whisky-calendar/domain"
)

type calendarDateController struct {
	service domain.CalendarDateUseCase
	ctx     context.Context
	router  router.Router
}

func NewWhiskyController(service domain.CalendarDateUseCase, ctx context.Context, router router.Router) domain.CalendarDateController {
	return &calendarDateController{
		service: service,
		ctx:     ctx,
		router:  router,
	}
}

func (c *calendarDateController) GetAllDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cdcd, err := c.service.Fetch(c.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("Couldn't fetch anything"))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cdcd)
}

func (c *calendarDateController) GetDateByDayOfMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	day, _ := c.router.ParsePathVariable(r, "day")
	dayInt, _ := strconv.ParseInt(day.(string), 10, 64)
	cd, _ := c.service.GetByDayOfMonth(c.ctx, dayInt)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cd)
}
