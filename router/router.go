package router

import (
	"goSqlx/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController *controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte("Hello World"))
	})

	router.POST("/book", bookController.CreateBook)
	router.GET("/book/:book_id", bookController.GetBookByID)

	return router
}
