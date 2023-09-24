package domain

import (
	"context"
	"net/http"
)

type Comment struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"userId"`
	WhiskyID int64  `json:"whiskyId"`
	Content  string `json:"content"`
	RegionID int64  `json:"regionId"`
	Rating   int    `json:"rating"`
}

type CommentUseCase interface {
	GetByID(ctx context.Context, id int64) (*Comment, error)
	GetByWhiskyID(ctx context.Context, whiskyId int64) ([]*Comment, error)
	Store(ctx context.Context, comment *Comment) error
}

type CommentStorer interface {
	Store(context.Context, *Comment) error
}

type CommentGetter interface {
	GetByID(context.Context, int64) (*Comment, error)
	GetByWhiskyID(context.Context, int64) ([]*Comment, error)
}

type CommentRepository interface {
	CommentStorer
	CommentGetter
}

type CommentController interface {
	GetCommentByID(w http.ResponseWriter, r *http.Request)
	GetCommentsByWhiskyID(w http.ResponseWriter, r *http.Request)
	StoreComment(w http.ResponseWriter, r *http.Request)
}

func NewComment(id int64, userId int64, whiskyId int64, content string, regionId int64, rating int) *Comment {
	return &Comment{
		ID:       id,
		UserID:   userId,
		WhiskyID: whiskyId,
		Content:  content,
		RegionID: regionId,
		Rating:   rating,
	}
}
