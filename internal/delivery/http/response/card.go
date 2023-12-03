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
	Cards         []entity.Card `json:"cards"`
	NumBlueCards  int           `json:"num_blue_cards"`
	NumRedCards   int           `json:"num_red_cards"`
	NumGreenCards int           `json:"num_green_cards"`
	User          *entity.User  `json:"user"`
}

type CopyCardToDeckResponse struct {
	Card entity.Card `json:"card"`
}
