package service

import (
	"context"
	"goSqlx/data/request"
	"goSqlx/data/response"
	"goSqlx/model"
	"goSqlx/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookRepositoryImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (b *BookServiceImpl) CreateBook(ctx context.Context, req *request.BookCreateRequest) error {

	err := b.BookRepository.SaveBook(ctx, &model.Book{
		BookID: req.BookID,
		Title:  req.Title,
		Price:  req.Price,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *BookServiceImpl) GetBookByID(ctx context.Context, bookID int64) (*response.BookResponse, error) {
	result, err := b.BookRepository.GetBookByID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, response.ErrBookNotFound
	}

	return result, nil
}
