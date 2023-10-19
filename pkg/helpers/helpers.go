package helpers

import (
	"time"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/pkg/timeutil"
)

func FilterReviewCards(rawCards *[]entity.Card, maxNewCards int, maxReviewCards int) *[]entity.Card {
	curTime := timeutil.TruncateToDay(time.Now())
	numNewCards := 0
	numRedCards := 0
	numReviewCards := 0
	var cards []entity.Card
	for _, card := range *rawCards {
		reviewTime := timeutil.TruncateToDay(card.NextReview)
		if reviewTime.After(curTime) {
			continue
		}
		if card.NumReviews == 0 {
			if numNewCards < maxNewCards {
				numNewCards++
				cards = append(cards, card)
			}
		} else if card.Sm2N == 0 {
			if numRedCards < maxReviewCards {
				numRedCards++
				cards = append(cards, card)
			}
		} else {
			if numReviewCards < maxReviewCards {
				numReviewCards++
				cards = append(cards, card)
			}
		}
	}
    return &cards
}
