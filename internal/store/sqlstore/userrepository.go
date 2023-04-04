package sqlstore

import (
	"smth/internal/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) CreateUser(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	_, err := r.FindByLogin(u.Email)

	if err == nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users(login, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
		u.Login,
		u.Email,
		u.PasswordHash,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, login, password_hash FROM users WHERE login = $1", login).Scan(&u.ID, &u.Login, &u.PasswordHash); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetRole() (int, error) {
	var roleId int
	if err := r.store.db.QueryRow("SELECT id FROM roles WHERE name = $1", "user").Scan(&roleId); err != nil {
		return 0, err
	}
	return roleId, nil
}
