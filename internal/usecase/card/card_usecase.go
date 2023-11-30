package card

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
	"vietcard-backend/pkg/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (uc *cardUsecase) GetReviewCardsByDeck(deckID *string, maxNewCards int, maxReviewCards int) (*[]entity.Card, int, int, int, error) {
	cards, err := uc.cardRepository.GetCardsByDeck(deckID)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	cards, numBlue, numRed, numGreen := helpers.FilterReviewCards(cards, maxNewCards, maxReviewCards)
	return cards, numBlue, numRed, numGreen, nil
}

func (uc *cardUsecase) UpdateCard(cardID *string, req *request.UpdateCardRequest) (*entity.Card, error) {
	return uc.cardRepository.UpdateCard(cardID, req)
}

func (uc *cardUsecase) UpdateCardReview(card *entity.Card) error {
	return uc.cardRepository.UpdateCardReview(card)
}

func (uc *cardUsecase) CopyCardToDeck(cardID *string, deckID *string) error {
	card, err := uc.cardRepository.GetCardByID(cardID)
	if err != nil {
		return err
	}
	card.ID = primitive.NilObjectID
	card.DeckID, err = primitive.ObjectIDFromHex(*deckID)
	card.SetDefault()
	if err != nil {
		return err
	}
	err = uc.cardRepository.CreateCard(card)
	if err != nil {
		return err
	}
	return nil
}
