package card

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
	"vietcard-backend/pkg/helpers"
)

type cardUsecase struct {
	cardRepository repository.CardRepository
	deckRepository repository.DeckRepository
}

func NewCardUsecase(cr repository.CardRepository, dr repository.DeckRepository) usecase.CardUsecase {
	return &cardUsecase{
		cardRepository: cr,
		deckRepository: dr,
	}
}

func (uc *cardUsecase) CreateCard(card *entity.Card) error {
	return uc.cardRepository.CreateCard(card)
}

func (uc *cardUsecase) GetCardByID(id *string) (*entity.Card, error) {
	return uc.cardRepository.GetCardByID(id)
}

func (uc *cardUsecase) GetReviewCardsByDeck(deckID *string, maxNewCards int, maxReviewCards int) (*[]entity.Card, error) {
	cards, err := uc.cardRepository.GetCardsByDeck(deckID)
	if err != nil {
		return nil, err
	}
	cards = helpers.FilterReviewCards(cards, maxNewCards, maxReviewCards)
	return cards, nil
}

func (uc *cardUsecase) UpdateCard(cardID *string, req *request.UpdateCardRequest) (*entity.Card, error) {
	return uc.cardRepository.UpdateCard(cardID, req)
}

func (uc *cardUsecase) UpdateCardReview(card *entity.Card) error {
	return uc.cardRepository.UpdateCardReview(card)
}
