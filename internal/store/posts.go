package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	ID       int64    `json:id`
	Content  string   `json:content`
	Title    string   `json:title`
	UserID   int64    `json:user_id`
	Tags     []string `json:tags`
	CreateAt string   `json:created_at`
	UpdateAt string   `json:updated_at`
}

type PostStore struct {
	db *sql.DB
}

func (p *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (content,tiltle, user_id,tags)
	VALUES ($1,$2,$3,$4) RETURNING id, created_at, updated_at`
	err := p.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreateAt,
		&post.UpdateAt,
	)
	if err != nil {
		return err
	}
	return nil
}
