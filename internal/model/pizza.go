package model

type Pizza struct {
	ID         int64   `json:"id" db:"id"`
	ImageURL   string  `json:"imageURL" db:"image_url"`
	Name       string  `json:"name" db:"name"`
	Types      []int64 `json:"types" db:"types"`
	Sizes      []int64 `json:"sizes" db:"sizes"`
	Price      int64   `json:"price" db:"price"`
	CategoryID int64   `json:"category_id" db:"category_id"`
	Rating     int64   `json:"rating" db:"rating"`
}
