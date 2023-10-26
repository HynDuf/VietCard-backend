package repository

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
)

type DeckRepository interface {
	CreateDeck(deck *entity.Deck) (*entity.Deck, error)
	GetDeckByID(id *string) (*entity.Deck, error)
	UpdateDeck(deckID *string, req *request.UpdateDeckRequest) (*entity.Deck, error)
	GetCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithReviewCards, error)
}
