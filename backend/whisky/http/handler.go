package http

import (
	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/gorilla/mux"
)

type whiskyHandler struct {
	wUseCase domain.WhiskyUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(mux *mux.Router, us domain.WhiskyUseCase) {
	handler := &ArticleHandler{
		AUsecase: us,
	}
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}
