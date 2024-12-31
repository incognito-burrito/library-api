package main

import (
	"log"
	"net/http"

	"github.com/incognito-burrito/library-api/books"
	"github.com/incognito-burrito/library-api/db"
	"github.com/incognito-burrito/library-api/repositories"
	"github.com/incognito-burrito/library-api/services"
)

func Run() error {
	dbConn := db.SetupSQLite("library.db")
	defer dbConn.Close()

	repos := repositories.InitializeRepos(dbConn)
	services := services.InitializeServices(repos)

	bookHandler := books.NewHandler(services.Book)

	mux := http.NewServeMux()
	mux.HandleFunc("/books", bookHandler.HandleBooks)
	mux.HandleFunc("/books/{id}", bookHandler.HandleSingleBook)

	log.Println("Server running on :8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
