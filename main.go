package main

import (
	"fmt"
	"goSqlx/config"
	"goSqlx/controller"
	"goSqlx/repository"
	"goSqlx/router"
	"goSqlx/service"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Start Server\n")

	// database
	db := config.DBConn()

	// repository
	bookRepo := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookRepositoryImpl(bookRepo)

	// controller
	bookController := controller.NewBookRepositoryImpl(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
