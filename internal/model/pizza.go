package model

type Pizza struct {
	ID         int    `json:"id" db:"id"`
	ImageURL   string `json:"imageURL" db:"image_url"`
	Name       string `json:"name" db:"name"`
	Types      []int  `json:"types" db:"types"`
	Sizes      []int  `json:"sizes" db:"sizes"`
	Price      int    `json:"price" db:"price"`
	CategoryID int    `json:"category_id" db:"category_id"`
	Rating     int    `json:"rating" db:"rating"`
}
