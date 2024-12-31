package repositories

import (
	"database/sql"

	"github.com/incognito-burrito/library-api/books"
)

type Repositories struct {
	Book *books.Repository
}

func InitializeRepos(db *sql.DB) *Repositories {
	return &Repositories{
		Book: books.NewRepository(db),
	}
}
