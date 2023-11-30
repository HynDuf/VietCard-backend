package helpers

import (
	"time"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/pkg/timeutil"
)

func FilterReviewCards(rawCards *[]entity.Card, maxNewCards int, maxReviewCards int) (*[]entity.Card, int, int, int) {
	curTime := timeutil.TruncateToDay(time.Now())
	numBlueCards := 0
	numRedCards := 0
	numGreenCards := 0
	var cards []entity.Card
	for _, card := range *rawCards {
		reviewTime := timeutil.TruncateToDay(card.NextReview)
		if reviewTime.After(curTime) {
			continue
		}
		if card.NumReviews == 0 {
			if numBlueCards < maxNewCards {
				numBlueCards++
				card.CardType = 0
				cards = append(cards, card)
			}
		} else if card.Sm2N == 0 {
			if numRedCards < maxReviewCards {
				numRedCards++
				card.CardType = 1
				cards = append(cards, card)
			}
		} else {
			if numGreenCards < maxReviewCards {
				numGreenCards++
				card.CardType = 2
				cards = append(cards, card)
			}
		}
	}
	return &cards, numBlueCards, numRedCards, numGreenCards
}
