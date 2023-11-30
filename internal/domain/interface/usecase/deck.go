package usecase

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
)

type DeckUsecase interface {
	CreateDeck(deck *entity.Deck) error
	GetDeckByID(id *string) (*entity.Deck, error)
	UpdateDeck(deckID *string, req *request.UpdateDeckRequest) (*entity.Deck, error)
	GetReviewCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithReviewCards, error)
	CopyDeck(userID *string, deckID *string) error
	GetDecksWithCards(userID *string) (*[]entity.DeckWithCards, *[]entity.DeckWithCards, *[]entity.DeckWithReviewCards, error)
}
