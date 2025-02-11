package books

type CreateBookDto struct {
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	CoverImageUrl string `json:"cover_image_url" validate:"required"`
}
