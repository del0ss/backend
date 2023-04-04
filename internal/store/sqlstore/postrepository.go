package sqlstore

import (
	"smth/internal/model"
)

type PostRepository struct {
	store *Store
}

func (r *PostRepository) CreatePost(p *model.Post, userId interface{}) error {
	return r.store.db.QueryRow(
		"INSERT INTO posts(user_id, title, content) VALUES ($1, $2, $3) RETURNING id",
		userId,
		p.Title,
		p.Content,
	).Scan(&p.ID)
}

func (r *PostRepository) GetPosts() ([]model.Post, error) {
	p := &model.Post{}
	var data []model.Post
	rows, err := r.store.db.Query("SELECT id, title, content FROM posts ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, *p)
	}
	return data, nil
}

func (r *PostRepository) GetPost(id int) (*model.Post, error) {
	p := &model.Post{}

	if err := r.store.db.QueryRow("SELECT * FROM posts WHERE id = $1", id).Scan(
		&p.ID,
		&p.UserID,
		&p.Title,
		&p.Content,
	); err != nil {
		return nil, err
	}
	return p, nil

}
