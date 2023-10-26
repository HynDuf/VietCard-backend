package repository

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
)

type CardRepository interface {
	CreateCard(card *entity.Card) error
    CreateManyCards(cards *[]entity.Card) error
	GetCardByID(id *string) (*entity.Card, error)
	GetCardsByDeck(deckID *string) (*[]entity.Card, error)
	UpdateCard(cardID *string, req *request.UpdateCardRequest) (*entity.Card, error)
	UpdateCardReview(card *entity.Card) error
}
