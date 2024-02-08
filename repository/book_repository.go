package repository

import (
	"context"
	"goSqlx/data/response"
	"goSqlx/model"
)

type BookRepository interface {
	SaveBook(ctx context.Context, book *model.Book) error
	GetBookByID(ctx context.Context, bookID int64) (*response.BookResponse, error)
	SearchBook()
}
