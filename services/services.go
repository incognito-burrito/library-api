package services

import (
	"github.com/incognito-burrito/library-api/books"
	"github.com/incognito-burrito/library-api/repositories"
)

type Services struct {
	Book *books.Service
}

func InitializeServices(repos *repositories.Repositories) *Services {
	return &Services{
		Book: books.NewService(repos.Book),
	}
}
