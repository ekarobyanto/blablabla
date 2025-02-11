package books

type Book struct {
	ID            int    `json:"id" db:"id"`
	Title         string `json:"title" db:"title"`
	Author        string `json:"author" db:"author"`
	CoverImageUrl string `json:"cover_image_url" db:"cover_image_url"`
	IsAvailable   bool   `json:"is_available" db:"is_available"`
}
