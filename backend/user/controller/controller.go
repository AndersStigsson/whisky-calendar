package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/AndersStigsson/whisky-calendar/token"
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

type userReturnModel struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
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

	u, err := c.service.Login(c.ctx, dest.TranslateToDomain())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errors.New("Incorrect username or password"))
		return
	}
	token, _ := token.GenerateJWTToken(u)
	userWithToken := userReturnModel{
		User:  u,
		Token: token,
	}
	json.NewEncoder(w).Encode(userWithToken)
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
	u, err := c.service.Register(c.ctx, dest.TranslateToDomain())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	token, _ := token.GenerateJWTToken(u)
	userWithToken := userReturnModel{
		User:  u,
		Token: token,
	}
	json.NewEncoder(w).Encode(userWithToken)
}

func (u *userBodyModel) TranslateToDomain() *domain.User {
	return domain.NewUser(
		u.ID,
		u.Username,
		u.Password,
		u.Name,
	)
}
