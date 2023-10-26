package usecase

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
)

type CardUsecase interface {
	CreateCard(card *entity.Card) error
	GetCardByID(id *string) (*entity.Card, error)
	GetReviewCardsByDeck(deckID *string, maxNewCards int, maxReviewCards int) (*[]entity.Card, error)
	UpdateCard(cardID *string, req *request.UpdateCardRequest) (*entity.Card, error)
	UpdateCardReview(card *entity.Card) error
    CopyCardToDeck(cardID *string, deckID *string) error
}
