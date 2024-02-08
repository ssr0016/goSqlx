package controller

import (
	"errors"
	"goSqlx/data/request"
	"goSqlx/helper"
	"goSqlx/model"
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

	if err := bookCreateRequest.Validate(); err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	err := c.BookService.CreateBook(r.Context(), &bookCreateRequest)
	if err != nil {
		if errors.Is(err, model.ErrBookTaken) {
			helper.WriteResponseBody(w, err.Error())
			return
		}
		helper.WriteResponseBody(w, err)
	}

	helper.WriteResponseBody(w, bookCreateRequest)
}

func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookID := ps.ByName("book_id")
	id, err := strconv.Atoi(bookID)
	if err != nil {

		helper.WriteResponseBody(w, err)
		return
	}

	result, err := c.BookService.GetBookByID(r.Context(), int64(id))
	if err != nil {
		helper.WriteResponseBody(w, err)
		return
	}

	helper.WriteResponseBody(w, result)
}

func (c *BookController) SearchBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	result, err := c.BookService.SearchBook(r.Context())
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	helper.WriteResponseBody(w, result)

}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(w, r, &bookUpdateRequest)

	bookID := ps.ByName("book_id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		helper.WriteResponseBody(w, err)
		return
	}

	if err := bookUpdateRequest.Validate(); err != nil {
		helper.WriteResponseBody(w, err.Error())
		return
	}

	bookUpdateRequest.ID = int64(id)
	err = c.BookService.UpdateBook(r.Context(), &bookUpdateRequest)
	if err != nil {
		helper.WriteResponseBody(w, err)
	}

	helper.WriteResponseBody(w, bookUpdateRequest)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookID := ps.ByName("book_id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		helper.WriteResponseBody(w, err)
		return
	}

	err = c.BookService.DeleteBook(r.Context(), int64(id))
	if err != nil {
		helper.WriteResponseBody(w, err)
	}

	helper.WriteResponseBody(w, nil)
}
