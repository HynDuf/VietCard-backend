package deck

import (
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
	"vietcard-backend/pkg/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deckUsecase struct {
	deckRepository repository.DeckRepository
	cardRepository repository.CardRepository
	userRepository repository.UserRepository
}

func NewDeckUsecase(dr repository.DeckRepository, cr repository.CardRepository, ur repository.UserRepository) usecase.DeckUsecase {
	return &deckUsecase{
		deckRepository: dr,
		cardRepository: cr,
		userRepository: ur,
	}
}

func (uc *deckUsecase) CreateDeck(deck *entity.Deck) error {
	_, err := uc.deckRepository.CreateDeck(deck)
	if err != nil {
		return err
	}
	return nil
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

func (uc *deckUsecase) CopyDeck(userID *string, deckID *string) error {
	user, err := uc.userRepository.GetByID(userID)
	if err != nil {
		return err
	}
	deck, err := uc.deckRepository.GetDeckByID(deckID)
	if err != nil {
		return err
	}
	cards, err := uc.cardRepository.GetCardsByDeck(deckID)
	if err != nil {
		return err
	}
	deck.ID = primitive.NilObjectID
	deck.UserID = user.ID
	if user.IsAdmin {
		deck.IsPublic = true
	} else {
		deck.IsPublic = false
	}
	if err != nil {
		return err
	}
	deck, err = uc.deckRepository.CreateDeck(deck)
	if err != nil {
		return err
	}
	for i := range *cards {
		(*cards)[i].UserID = deck.UserID
		(*cards)[i].DeckID = deck.ID
		(*cards)[i].ID = primitive.NilObjectID
	}
	err = uc.cardRepository.CreateManyCards(cards)
	if err != nil {
		return err
	}
	return nil
}
