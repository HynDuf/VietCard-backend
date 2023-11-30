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
	decksWithReviewCards := []entity.DeckWithReviewCards{}
	for i := range *rawDeckWithCards {
		rawDeck := (*rawDeckWithCards)[i]
		deck := entity.DeckWithReviewCards{
			Deck:          rawDeck.Deck, // Copy fields from the embedded Deck
			Cards:         rawDeck.Cards,
			NumBlueCards:  0, // Set the initial values for NumBlueCards, NumRedCards, NumGreenCards
			NumRedCards:   0,
			NumGreenCards: 0,
		}
		deck.UpdateReview()
		deck.Cards, deck.NumBlueCards, deck.NumRedCards, deck.NumGreenCards = helpers.FilterReviewCards(deck.Cards, deck.MaxNewCards-deck.CurNewCards, deck.MaxReviewCards-deck.CurReviewCards)
		decksWithReviewCards = append(decksWithReviewCards, deck)
	}
	return &decksWithReviewCards, nil
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

func (uc *deckUsecase) GetDecksWithCards(userID *string) (*[]entity.DeckWithCards, *[]entity.DeckWithCards, *[]entity.DeckWithReviewCards, error) {
	rawDeckWithCards, err := uc.deckRepository.GetCardsAllDecks(userID)
	if err != nil {
		return nil, nil, nil, err
	}
	userDecks := []entity.DeckWithCards{}
	publicDecks := []entity.DeckWithCards{}
	for _, deck := range *rawDeckWithCards {
		if deck.IsPublic {
			publicDecks = append(publicDecks, deck)
		} else {
			userDecks = append(userDecks, deck)
		}
	}
	decksWithReviewCards := []entity.DeckWithReviewCards{}
	for i := range userDecks {
		rawDeck := userDecks[i]
		deck := entity.DeckWithReviewCards{
			Deck:          rawDeck.Deck, // Copy fields from the embedded Deck
			Cards:         rawDeck.Cards,
			NumBlueCards:  0, // Set the initial values for NumBlueCards, NumRedCards, NumGreenCards
			NumRedCards:   0,
			NumGreenCards: 0,
		}
		deck.UpdateReview()
		deck.Cards, deck.NumBlueCards, deck.NumRedCards, deck.NumGreenCards = helpers.FilterReviewCards(deck.Cards, deck.MaxNewCards-deck.CurNewCards, deck.MaxReviewCards-deck.CurReviewCards)
		decksWithReviewCards = append(decksWithReviewCards, deck)
	}
	return &userDecks, &publicDecks, &decksWithReviewCards, nil
}
