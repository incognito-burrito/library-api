package books

import (
	"database/sql"

	"github.com/incognito-burrito/library-api/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *Repository) Create(book models.Book) (int, error) {
	result, err := r.db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *Repository) GetByID(id int) (models.Book, error) {
	var book models.Book
	err := r.db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, nil
		}
		return book, err
	}
	return book, nil
}

func (r *Repository) Update(book models.Book) error {
	_, err := r.db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", book.Title, book.Author, book.ID)
	return err
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
