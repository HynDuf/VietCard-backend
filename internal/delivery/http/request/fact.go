package request

type CreateFactRequest struct {
	Content string `json:"content" binding:"required"`
}
