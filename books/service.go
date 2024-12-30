package books

import (
	"github.com/incognito-burrito/library-api/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *Service) AddBook(book models.Book) (models.Book, error) {
	id, err := s.repo.Create(book)
	if err != nil {
		return models.Book{}, err
	}
	book.ID = id
	return book, nil
}

func (s *Service) GetBook(id int) (models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdateBook(id int, book models.Book) (models.Book, error) {
	book.ID = id
	err := s.repo.Update(book)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s *Service) DeleteBook(id int) error {
	return s.repo.Delete(id)
}
