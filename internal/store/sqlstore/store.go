package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"smth/internal/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	postRepository *PostRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

func (s *Store) Post() store.PostRepository {
	if s.postRepository != nil {
		return s.postRepository
	}
	s.postRepository = &PostRepository{
		store: s,
	}
	return s.postRepository
}
