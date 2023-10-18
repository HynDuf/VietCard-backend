package deck

import (
	"time"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"
	"vietcard-backend/internal/domain/interface/usecase"
	"vietcard-backend/pkg/timeutil"
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

func (uc *deckUsecase) UpdateDeck(deck *entity.Deck) error {
	return uc.deckRepository.UpdateDeck(deck)
}

func (uc *deckUsecase) GetReviewCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithReviewCards, error) {
	rawDeckWithCards, err := uc.deckRepository.GetCardsAllDecksOfUser(userID)
	if err != nil {
		return nil, err
	}
	user, err := uc.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	numNewCards := 0
	numRedCards := 0
	numReviewCards := 0
    curTime := timeutil.TruncateToDay(time.Now())
	for i := range *rawDeckWithCards {
		deck := (*rawDeckWithCards)[i]
		var cards []entity.Card
		for _, card := range deck.Cards {
			reviewTime := timeutil.TruncateToDay(card.NextReview)
            if reviewTime.After(curTime) {
                continue
            }
			if card.NumReviews == 0 {
				if numNewCards < user.MaxNewCardsLearn {
					numNewCards++
					cards = append(cards, card)
				}
			} else if card.Sm2N == 0 {
				if numRedCards < user.MaxCardsReview {
					numRedCards++
					cards = append(cards, card)
				}
			} else {
				if numReviewCards < user.MaxCardsReview {
					numReviewCards++
					cards = append(cards, card)
				}
			}
		}
		deck.Cards = cards
		(*rawDeckWithCards)[i] = deck
	}
	return rawDeckWithCards, nil
}
