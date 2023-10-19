package deck

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
	"vietcard-backend/pkg/helpers"
)

type deckUsecase struct {
	deckRepository repository.DeckRepository
	userRepository repository.UserRepository
}

func NewDeckUsecase(dr repository.DeckRepository, ur repository.UserRepository) usecase.DeckUsecase {
	return &deckUsecase{
		deckRepository: dr,
		userRepository: ur,
	}
}

func (uc *deckUsecase) CreateDeck(deck *entity.Deck) error {
	return uc.deckRepository.CreateDeck(deck)
}

func (uc *deckUsecase) GetDeckByID(id *string) (*entity.Deck, error) {
	return uc.deckRepository.GetDeckByID(id)
}

func (uc *deckUsecase) UpdateDeck(deckID *string, req *request.UpdateDeckRequest) (*entity.Deck, error) {
	return uc.deckRepository.UpdateDeck(deckID, req)
}

func (uc *deckUsecase) GetReviewCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithReviewCards, error) {
	rawDeckWithCards, err := uc.deckRepository.GetCardsAllDecksOfUser(userID)
	if err != nil {
		return nil, err
	}
	for i := range *rawDeckWithCards {
		deck := (*rawDeckWithCards)[i]
		deck.UpdateReview()
		deck.Cards = helpers.FilterReviewCards(deck.Cards, deck.MaxNewCards-deck.CurNewCards, deck.MaxReviewCards-deck.CurReviewCards)
		(*rawDeckWithCards)[i] = deck
	}
	return rawDeckWithCards, nil
}
