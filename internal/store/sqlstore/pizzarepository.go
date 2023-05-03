package sqlstore

import (
	"fmt"
	"github.com/lib/pq"
	"smth/internal/model"
)

type PizzaRepository struct {
	store *Store
}

func (r *PizzaRepository) CreatePizza(p model.Pizza) (int, error) {
	var pizzaID int
	tx, err := r.store.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	row := r.store.db.QueryRow("INSERT INTO pizzas(image_url, name, types, sizes, price, category_id, rating) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		p.ImageURL,
		p.Name,
		pq.Array(p.Types),
		pq.Array(p.Sizes),
		p.Price,
		p.CategoryID,
		p.Rating,
	)

	errScan := row.Scan(&pizzaID)
	if errScan != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return pizzaID, nil
}

func (r *PizzaRepository) GetPizza(sort string) ([]model.Pizza, error) {
	p := &model.Pizza{}
	var data []model.Pizza
	var query string
	if sort != "" {
		query = fmt.Sprintf("SELECT * FROM pizzas ORDER BY %s", sort)
	} else {
		query = fmt.Sprintf("SELECT * FROM pizzas")
	}
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&p.ID,
			&p.ImageURL,
			&p.Name,
			pq.Array(&p.Types),
			pq.Array(&p.Sizes),
			&p.Price,
			&p.CategoryID,
			&p.Rating,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, *p)
	}
	return data, nil
}

func (r *PizzaRepository) GetCategories() ([]model.Category, error) {
	p := &model.Category{}
	var data []model.Category
	rows, err := r.store.db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&p.ID,
			&p.Name,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, *p)
	}
	return data, nil
}

func (r *PizzaRepository) GetCategoryById(id int) (*model.Category, error) {
	p := &model.Category{}
	if err := r.store.db.QueryRow("SELECT * FROM category WHERE id = $1", id).Scan(
		&p.ID,
		&p.Name,
	); err != nil {
		return nil, err
	}
	return p, nil

}

func (r *PizzaRepository) GetPizzaById(id int) (*model.Pizza, error) {
	p := &model.Pizza{}

	if err := r.store.db.QueryRow("SELECT * FROM pizzas WHERE id = $1", id).Scan(
		&p.ID,
		&p.ImageURL,
		&p.Name,
		&p.Types,
		&p.Sizes,
		&p.Price,
		&p.CategoryID,
		&p.Rating,
	); err != nil {
		return nil, err
	}
	return p, nil

}
func (r *PizzaRepository) DeletePizza(id int) error {
	_, err := r.store.db.Exec("DELETE FROM pizzas WHERE id = $1", id)
	return err
}
