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
	result, err := b.BookRepository.BookTaken(ctx, 0, req.Title)
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return model.ErrBookTaken
	}
	err = b.BookRepository.SaveBook(ctx, &model.Book{
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
func (b *BookServiceImpl) SearchBook(ctx context.Context) ([]*response.BookResponse, error) {

	results, err := b.BookRepository.SearchBook(ctx)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (b *BookServiceImpl) UpdateBook(ctx context.Context, req *request.BookUpdateRequest) error {

	err := b.BookRepository.UpdateBook(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookServiceImpl) DeleteBook(ctx context.Context, bookID int64) error {
	result, err := b.BookRepository.GetBookByID(ctx, bookID)
	if err != nil {
		return err
	}

	if result == nil {
		return response.ErrBookNotFound
	}

	err = b.BookRepository.DeleteBook(ctx, bookID)
	if err != nil {
		return err
	}

	return nil
}
