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

func (s *Service) All() ([]models.Book, error) {
	return s.repo.All()
}

func (s *Service) Add(book models.Book) (models.Book, error) {
	id, err := s.repo.Create(book)
	if err != nil {
		return models.Book{}, err
	}
	book.ID = id
	return book, nil
}

func (s *Service) Get(id int) (models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id int, book models.Book) (models.Book, error) {
	book.ID = id
	err := s.repo.Update(book)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
