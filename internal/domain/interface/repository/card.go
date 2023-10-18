package repository

import (
	"vietcard-backend/internal/domain/entity"
)

type CardRepository interface {
	CreateCard(card *entity.Card) error
	GetCardByID(id *string) (*entity.Card, error)
    UpdateCard(card *entity.Card) error
    UpdateCardReview(card *entity.Card) error
}
