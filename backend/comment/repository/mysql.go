package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AndersStigsson/whisky-calendar/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type commentRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

type CommentModel struct {
	ID       int64  `db:"id"`
	UserID   int64  `db:"user_id"`
	WhiskyID int64  `db:"whisky_id"`
	Content  string `db:"content"`
	RegionID int64  `db:"region_id"`
	Rating   int    `db:"rating"`
}

func (r *commentRepository) GetByID(ctx context.Context, id int64) (*domain.Comment, error) {
	cm := &CommentModel{}
	if err := r.db.GetContext(ctx, cm, getCommentByID, id); err != nil {
		return nil, errors.Wrap(err, "regionRepo.GetCommentByID.GetContext")
	}

	return cm.TranslateToDomain()
}

func (r *commentRepository) GetByWhiskyID(ctx context.Context, whiskyId int64) ([]*domain.Comment, error) {
	var cc []*CommentModel
	rows, err := r.db.QueryxContext(ctx, getCommentsByWhiskyID, whiskyId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "commentRepo.Fetch.QueryxContext")
		}
		return []*domain.Comment{}, nil
	}

	defer rows.Close()
	for rows.Next() {
		c := &CommentModel{}
		if err = rows.StructScan(c); err != nil {
			return nil, errors.Wrap(err, "whiskyRepo.Fetch.Structscan")
		}
		cc = append(cc, c)
	}

	var dcc []*domain.Comment
	for _, c := range cc {
		dc, err := c.TranslateToDomain()
		if err != nil {
			fmt.Printf("Do something with this: %v", err)
		}
		dcc = append(dcc, dc)
	}

	return dcc, nil
}

func (r *commentRepository) Store(ctx context.Context, comment *domain.Comment) error {
	c := &CommentModel{}
	if err := r.db.QueryRowxContext(
		ctx,
		storeComment,
		comment.UserID,
		comment.WhiskyID,
		comment.Content,
		comment.RegionID,
		comment.Rating,
	).StructScan(c); err != nil {
		fmt.Printf("Err: %v\n", err.Error())
		return err
	}
	return nil
}

func (cm *CommentModel) TranslateToDomain() (*domain.Comment, error) {
	dc := domain.NewComment(
		cm.ID,
		cm.UserID,
		cm.WhiskyID,
		cm.Content,
		cm.RegionID,
		cm.Rating,
	)
	return dc, nil
}
