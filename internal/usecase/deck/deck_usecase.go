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

func (uc *deckUsecase) CreateDeck(deck *entity.Deck) (*entity.Deck, error) {
	deck, err := uc.deckRepository.CreateDeck(deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
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

func (uc *deckUsecase) CopyDeck(userID *string, deckID *string) (*entity.DeckWithCards, *entity.DeckWithReviewCards, error) {
	user, err := uc.userRepository.GetByID(userID)
	if err != nil {
		return nil, nil, err
	}
	deck, err := uc.deckRepository.GetDeckByID(deckID)
	if err != nil {
		return nil, nil, err
	}
	cards, err := uc.cardRepository.GetCardsByDeck(deckID)
	if err != nil {
		return nil, nil, err
	}
	deck.ID = primitive.NilObjectID
	deck.UserID = user.ID
	deck.IsPublic = false
	if err != nil {
		return nil, nil, err
	}
	deck, err = uc.deckRepository.CreateDeck(deck)
	if err != nil {
		return nil, nil, err
	}
	for i := range *cards {
		(*cards)[i].UserID = deck.UserID
		(*cards)[i].DeckID = deck.ID
		(*cards)[i].ID = primitive.NilObjectID
	}
	err = uc.cardRepository.CreateManyCards(cards)
	if err != nil {
		return nil, nil, err
	}
	dID := deck.ID.Hex()
	rawDeckWithCards, err := uc.deckRepository.GetDeckWithCards(&dID)
	if err != nil {
		return nil, nil, err
	}
	deckWithReviewCards := entity.DeckWithReviewCards{
		Deck:          rawDeckWithCards.Deck,
		Cards:         rawDeckWithCards.Cards,
		NumBlueCards:  0,
		NumRedCards:   0,
		NumGreenCards: 0,
	}
	deckWithReviewCards.UpdateReview()
	deckWithReviewCards.Cards, deckWithReviewCards.NumBlueCards, deckWithReviewCards.NumRedCards, deckWithReviewCards.NumGreenCards = helpers.FilterReviewCards(deckWithReviewCards.Cards, deckWithReviewCards.MaxNewCards-deckWithReviewCards.CurNewCards, deckWithReviewCards.MaxReviewCards-deckWithReviewCards.CurReviewCards)
	rawDeckWithCards.NumBlueCards = deckWithReviewCards.NumBlueCards
	rawDeckWithCards.NumRedCards = deckWithReviewCards.NumRedCards
	rawDeckWithCards.NumGreenCards = deckWithReviewCards.NumGreenCards
	return rawDeckWithCards, &deckWithReviewCards, nil
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
		userDecks[i].NumBlueCards = deck.NumBlueCards
		userDecks[i].NumRedCards = deck.NumRedCards
		userDecks[i].NumGreenCards = deck.NumGreenCards
	}
	return &userDecks, &publicDecks, &decksWithReviewCards, nil
}
func (uc *deckUsecase) DeleteDeck(deckID *string) error {
	return uc.deckRepository.DeleteDeck(deckID)
}
