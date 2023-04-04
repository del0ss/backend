package model

type Post struct {
	ID      int    `json:"-" db:"id"`
	UserID  int    `json:"-"`
	Title   string `json:"title" binding:"required,min=4,max=15"`
	Content string `json:"content" binding:"required"`
}
