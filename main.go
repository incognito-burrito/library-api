package main

import (
	"log"
	"net/http"

	"github.com/incognito-burrito/library-api/books"
	"github.com/incognito-burrito/library-api/db"
)

func main() {
	dbConn := db.SetupSQLite("library.db")
	defer dbConn.Close()

	bookRepo := books.NewRepository(dbConn)
	bookService := books.NewService(bookRepo)
	bookHandler := books.NewHandler(bookService)

	mux := http.NewServeMux()

	mux.HandleFunc("/books", bookHandler.HandleBooks)
	mux.HandleFunc("/books/{id}", bookHandler.HandleSingleBook)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
