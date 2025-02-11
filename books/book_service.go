package books

type BookService struct {
	repo *BookRepo
}

func NewBookService(repo *BookRepo) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllAvailableBooks() ([]Book, error) {
	return s.repo.FindAll()
}

func (s *BookService) FindByID(id int) (*Book, error) {
	return s.repo.FindByID(id)
}

func (s *BookService) FindByTitle(title string) ([]Book, error) {
	return s.repo.FindByTitle(title)
}

func (s *BookService) Create(book *CreateBookDto) error {
	createBook := Book{
		Title:         book.Title,
		Author:        book.Author,
		CoverImageUrl: book.CoverImageUrl,
	}
	return s.repo.Create(&createBook)
}
