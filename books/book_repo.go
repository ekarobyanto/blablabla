package books

import "github.com/jmoiron/sqlx"

type BookRepo struct {
	con *sqlx.DB
}

func NewBookRepo(con *sqlx.DB) *BookRepo {
	return &BookRepo{con: con}
}

func (r *BookRepo) FindAll() ([]Book, error) {
	var books []Book
	err := r.con.Select(&books, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) FindByID(id int) (*Book, error) {
	book := Book{}
	err := r.con.Get(&book, "SELECT * FROM books WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepo) FindByTitle(title string) ([]Book, error) {
	var books []Book
	err := r.con.Select(&books, "SELECT * FROM books WHERE title LIKE $1", "%"+title+"%")
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepo) Create(book *Book) error {
	_, err := r.con.Exec("INSERT INTO books (title, author, cover_image_url, is_available) VALUES ($1, $2, $3)", book.Title, book.Author, true)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepo) Update(book *Book) error {
	_, err := r.con.Exec("UPDATE books SET title = coalesce($1, title), author = coalesce($2, author), cover_image_url = coalesce($3, cover_image_url), is_available = coalesce($4, is_available) WHERE id = $5", book.Title, book.Author, book.CoverImageUrl, book.IsAvailable, book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepo) GetAllAvailableBooks() ([]Book, error) {
	var books []Book
	err := r.con.Select(&books, "Select * From books WHERE is_available = true")
	if err != nil {
		return []Book{}, err
	}
	return books, nil
}
