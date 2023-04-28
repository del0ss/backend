package model

type Post struct {
	ID      int64  `json:"id" db:"id"`
	UserID  int64  `json:"user_id"`
	Title   string `json:"title" binding:"required,min=4,max=15"`
	Content string `json:"content" binding:"required"`
}
