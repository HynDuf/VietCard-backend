package usecase

import "vietcard-backend/internal/domain/entity"

type DeckUsecase interface {
	CreateDeck(deck *entity.Deck) error
	GetDeckByID(id *string) (*entity.Deck, error)
	UpdateDeck(deck *entity.Deck) error
	GetReviewCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithReviewCards, error)
}
