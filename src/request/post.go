package request

type (
	CreatePostRequest struct {
		Title      string `json:"title" validate:"required"`
		Content    string `json:"content" validate:"required"`
		CategoryId uint   `json:"category_id" validate:"required"`
	}

	UpdatePostRequest struct {
		Title      string `json:"title" validate:"required_without_all=Content CategoryId"`
		Content    string `json:"content" validate:"required_without_all=Title CategoryId"`
		CategoryId uint   `json:"category_id" validate:"required_without_all=Title Content"`
	}
)
