package response

import "vietcard-backend/internal/domain/entity"

type CreateCardResponse struct {
	Card entity.Card `json:"card"`
}

type UpdateCardResponse struct {
	Success bool        `json:"success"`
	Card    entity.Card `json:"card"`
}

type UpdateReviewCardsResponse struct {
	Success       bool          `json:"success"`
	Cards         []entity.Card `json:"cards"`
	NumBlueCards  int           `json:"num_blue_cards"`
	NumRedCards   int           `json:"num_red_cards"`
	NumGreenCards int           `json:"num_green_cards"`
}

type CopyCardToDeckResponse struct {
	Card entity.Card `json:"card"`
}
