package repository

import (
	"context"
	"goSqlx/data/request"
	"goSqlx/data/response"
	"goSqlx/model"
)

type BookRepository interface {
	SaveBook(ctx context.Context, book *model.Book) error
	GetBookByID(ctx context.Context, bookID int64) (*response.BookResponse, error)
	SearchBook(ctx context.Context) ([]*response.BookResponse, error)
	UpdateBook(ctx context.Context, req *request.BookUpdateRequest) error
	DeleteBook(ctx context.Context, bookID int64) error

	BookTaken(ctx context.Context, id int64, title string) ([]*model.Book, error)
}
