package response

import (
	"errors"
)

var (
	ErrBookNotFound   = errors.New("book not found")
	ErrInvalidRequest = errors.New("invalid request")
)

type BookResponse struct {
	ID     int64   `db:"id" json:"id"`
	BookID string  `db:"book_id" json:"book_id"`
	Title  string  `db:"title" json:"title"`
	Price  float64 `db:"price" json:"price"`
}
