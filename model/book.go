package model

import "errors"

var (
	ErrBookIDEmpty = errors.New("book id cannot be empty")
	ErrTitleEmpty  = errors.New("title cannot be empty")
	ErrPriceEmpty  = errors.New("price cannot be empty")
	ErrBookTaken   = errors.New("book already taken")
)

type Book struct {
	ID     int64   `db:"id" json:"id"`
	BookID string  `db:"book_id" json:"book_id"`
	Title  string  `db:"title" json:"title"`
	Price  float64 `db:"price" json:"price"`
}
