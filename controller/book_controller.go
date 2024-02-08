package controller

import (
	"goSqlx/data/request"
	"goSqlx/helper"
	"goSqlx/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct {
	BookService service.BookService
}

func NewBookRepositoryImpl(bookService service.BookService) *BookController {
	return &BookController{
		BookService: bookService,
	}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(w, r, &bookCreateRequest)

	err := c.BookService.CreateBook(r.Context(), &bookCreateRequest)
	if err != nil {
		helper.WriteResponseBody(w, err)
	}

	helper.WriteResponseBody(w, bookCreateRequest)
}

func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookID := ps.ByName("book_id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		// If strconv.Atoi() fails, return an error response
		helper.WriteResponseBody(w, err)
		return // Exit the handler function
	}

	// Call the service method to get the book by ID
	result, err := c.BookService.GetBookByID(r.Context(), int64(id))
	if err != nil {
		// If an error occurs, return an error response
		helper.WriteResponseBody(w, err)
		return // Exit the handler function
	}

	// Write the book response to the response writer
	helper.WriteResponseBody(w, result)
}
