package service

import (
	"context"
	"goSqlx/data/request"
	"goSqlx/data/response"
)

type BookService interface {
	CreateBook(ctx context.Context, req *request.BookCreateRequest) error
	GetBookByID(ctx context.Context, bookID int64) (*response.BookResponse, error)
	SearchBook(ctx context.Context) ([]*response.BookResponse, error)
	UpdateBook(ctx context.Context, req *request.BookUpdateRequest) error
	DeleteBook(ctx context.Context, bookID int64) error
}
