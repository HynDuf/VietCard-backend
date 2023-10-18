package deck

import (
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
)

type deckUsecase struct {
	deckRepository repository.DeckRepository
}

func NewDeckUsecase(dr repository.DeckRepository) usecase.DeckUsecase {
	return &deckUsecase{
		deckRepository: dr,
	}
}
func (uc *deckUsecase) CreateDeck(deck *entity.Deck) error {
	return uc.deckRepository.CreateDeck(deck)
}
func (uc *deckUsecase) GetDeckByID(id *string) (*entity.Deck, error) {
	return uc.deckRepository.GetDeckByID(id)
}
func (uc *deckUsecase) UpdateDeck(deck *entity.Deck) error {
	return uc.deckRepository.UpdateDeck(deck)
}
func (uc *deckUsecase) GetReviewCardsAllDecksOfUser(userID *string) (*entity.DeckWithReviewCards, error) {
	return uc.deckRepository.GetReviewCardsAllDecksOfUser(userID)
}
