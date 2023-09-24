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

type commentController struct {
	service domain.CommentUseCase
	ctx     context.Context
	router  router.Router
}

func New(service domain.CommentUseCase, ctx context.Context, router router.Router) domain.CommentController {
	return &commentController{
		service: service,
		ctx:     ctx,
		router:  router,
	}
}

type commentBodyModel struct {
	UserID   int64  `json:"userId"`
	WhiskyID int64  `json:"whiskyId"`
	Content  string `json:"content"`
	RegionID int64  `json:"regionId"`
	Rating   int    `json:"rating"`
}

func (c *commentController) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := c.router.ParsePathVariable(r, "id")
	idint, _ := strconv.ParseInt(id.(string), 10, 64)
	res, _ := c.service.GetByID(c.ctx, idint)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (c *commentController) GetCommentsByWhiskyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := c.router.ParsePathVariable(r, "whiskyId")
	idint, _ := strconv.ParseInt(id.(string), 10, 64)
	res, _ := c.service.GetByWhiskyID(c.ctx, idint)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (c *commentController) StoreComment(w http.ResponseWriter, r *http.Request) {
	var dest commentBodyModel
	bodyBytes, err := c.router.GetBody(r)
	json.Unmarshal(bodyBytes, &dest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.New("Unable to get body"))
	}
	c.service.Store(c.ctx, dest.TranslateToDomain())
}

func (c *commentBodyModel) TranslateToDomain() *domain.Comment {
	return domain.NewComment(
		0,
		c.UserID,
		c.WhiskyID,
		c.Content,
		c.RegionID,
		c.Rating,
	)
}
