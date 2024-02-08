package repository

import (
	"context"
	"database/sql"
	"errors"
	"goSqlx/data/request"
	"goSqlx/data/response"
	"goSqlx/model"

	"github.com/jmoiron/sqlx"
)

type BookRepositoryImpl struct {
	Db *sqlx.DB
}

func NewBookRepository(Db *sqlx.DB) BookRepository {
	return &BookRepositoryImpl{
		Db: Db,
	}
}

func (b *BookRepositoryImpl) SaveBook(ctx context.Context, book *model.Book) error {
	tx := b.Db.MustBeginTx(ctx, nil)
	defer tx.Rollback()

	rawSQL := `
		INSERT INTO books (book_id, title, price)
		VALUES (:book_id, :title, :price)
	`

	_, err := tx.NamedExecContext(ctx, rawSQL, book)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepositoryImpl) GetBookByID(ctx context.Context, bookID int64) (*response.BookResponse, error) {
	var result response.BookResponse

	rawSQL := `
		SELECT
			id,
			book_id,
			title,
			price
		FROM
			books
		WHERE
			id = $1
	`

	err := b.Db.GetContext(ctx, &result, rawSQL, bookID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &result, nil
}

func (b *BookRepositoryImpl) SearchBook(ctx context.Context) ([]*response.BookResponse, error) {
	var results []*response.BookResponse

	rawSQL := `
        SELECT
            id,
            book_id,
            title,
            price
        FROM
            books
    `

	err := b.Db.SelectContext(ctx, &results, rawSQL)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (b *BookRepositoryImpl) UpdateBook(ctx context.Context, req *request.BookUpdateRequest) error {
	tx, err := b.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rawSQL := `
			UPDATE
				books
			SET
				book_id = :book_id,
				title = :title,
				price = :price
			WHERE
				id = :id
		`

	_, err = b.Db.NamedExecContext(ctx, rawSQL, req)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepositoryImpl) DeleteBook(ctx context.Context, bookID int64) error {
	tx, err := b.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rawSQL := `
		DELETE FROM
			books
		WHERE
			id = $1
	`

	_, err = tx.ExecContext(ctx, rawSQL, bookID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepositoryImpl) BookTaken(ctx context.Context, id int64, title string) ([]*model.Book, error) {
	var result []*model.Book

	rawSQL := `
        SELECT
            id,
            book_id,
            title,
            price
        FROM
            books
        WHERE
            book_id = $1 OR
            title = $2
    `
	err := b.Db.SelectContext(ctx, &result, rawSQL, id, title)
	if err != nil {
		return nil, err
	}

	return result, nil
}
