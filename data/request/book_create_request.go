package request

import "goSqlx/model"

type BookCreateRequest struct {
	BookID string  `db:"book_id" json:"book_id"`
	Title  string  `db:"title" json:"title"`
	Price  float64 `db:"price" json:"price"`
}

func (b *BookCreateRequest) Validate() error {
	if len(b.BookID) <= 0 {
		return model.ErrBookIDEmpty
	}

	if len(b.Title) == 0 {
		return model.ErrTitleEmpty
	}

	if b.Price <= 0 || b.Price == 0 {
		return model.ErrPriceEmpty
	}

	return nil
}
