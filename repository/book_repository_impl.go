package repository

import (
	"context"
	"database/sql"
	"errors"
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
