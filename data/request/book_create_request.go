package request

type BookCreateRequest struct {
	BookID string  `db:"book_id" json:"book_id"`
	Title  string  `db:"title" json:"title"`
	Price  float64 `db:"price" json:"price"`
}
