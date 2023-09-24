package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	"github.com/AndersStigsson/whisky-calendar/domain"
)

type userController struct {
	service domain.UserUseCase
	ctx     context.Context
	router  router.Router
}

func New(service domain.UserUseCase, ctx context.Context, router router.Router) domain.UserController {
	return &userController{
		service: service,
		ctx:     ctx,
		router:  router,
	}
}

type userBodyModel struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name,omitempty"`
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	var dest userBodyModel
	bodyBytes, err := c.router.GetBody(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("Unable to get body"))
		return
	}
	err = json.Unmarshal(bodyBytes, &dest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("Unable to marshal body"))
		return
	}

	err = c.service.Login(c.ctx, dest.TranslateToDomain())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errors.New("Incorrect username or password"))
		return
	}
	json.NewEncoder(w).Encode("Login Succeeded")
}

func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	var dest userBodyModel
	bodyBytes, err := c.router.GetBody(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("Unable to get body"))
		return
	}
	err = json.Unmarshal(bodyBytes, &dest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("Unable to unmarshal body"))
		return
	}
	err = c.service.Register(c.ctx, dest.TranslateToDomain())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode("Registered")
}

func (u *userBodyModel) TranslateToDomain() *domain.User {
	return domain.NewUser(
		u.ID,
		u.Username,
		u.Password,
		u.Name,
	)
}
