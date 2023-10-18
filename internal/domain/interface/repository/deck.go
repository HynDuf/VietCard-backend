package repository

import (
	"vietcard-backend/internal/domain/entity"
)

type DeckRepository interface {
	CreateDeck(deck *entity.Deck) error
	GetDeckByID(id *string) (*entity.Deck, error)
	UpdateDeck(deck *entity.Deck) error
	GetReviewCardsAllDecksOfUser(userID *string) (*entity.DeckWithReviewCards, error)
}
