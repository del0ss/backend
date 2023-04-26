package store

import (
	"smth/internal/model"
)

type UserRepository interface {
	CreateUser(*model.User) error
	FindByLogin(login string) (model.User, error)
	GetRole() (int, error)
}

type PostRepository interface {
	CreatePost(post model.Post, userId interface{}) (int, error)
	GetPosts() ([]model.Post, error)
	GetPost(id int) (*model.Post, error)
	DeletePost(id int) error
}
