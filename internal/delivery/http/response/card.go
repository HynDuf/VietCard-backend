package response

import "vietcard-backend/internal/domain/entity"

type CreateCardResponse struct {
	Success bool `json:"success"`
}

type UpdateCardResponse struct {
	Success bool        `json:"success"`
	Card    entity.Card `json:"card"`
}
