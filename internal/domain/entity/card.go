package entity

import (
	"math"
	"math/rand"
	"time"
	"vietcard-backend/pkg/timeutil"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	DeckID       primitive.ObjectID `json:"deck_id" bson:"deck_id"`
	Question     string             `json:"question" bson:"question"`
	Answer       string             `json:"answer" bson:"answer"`
	WrongAnswers []string           `json:"wrong_answers" bson:"wrong_answers"`
	LastReview   time.Time          `json:"last_review" bson:"last_review"`
	NextReview   time.Time          `json:"next_review" bson:"next_review"`
	NumReviews   int                `json:"num_reviews" bson:"num_reviews"`
	Sm2N         int                `json:"sm2_n" bson:"sm2_n"`
	Sm2EF        float64            `json:"sm2_ef" bson:"sm2_ef"`
	Sm2I         int                `json:"sm2_i" bson:"sm2_i"`
}

func (card *Card) SetDefault() *Card {
	card.CreatedAt = time.Now()
	card.NextReview = card.CreatedAt
	card.NumReviews = 0
	card.Sm2N = 0
	card.Sm2EF = 2.5
	card.Sm2I = 0
	return card
}

/*
 * Reference: https://en.wikipedia.org/wiki/SuperMemo#Description_of_SM-2_algorithm
 */
func (card *Card) UpdateScheduleSM2(correct bool) *Card {
	n := card.Sm2N
	EF := card.Sm2EF
	I := card.Sm2I
	oldI := I
	if correct {
		if n == 0 {
			I = 1
		} else if n == 1 {
			I = 4
		} else {
			I = int(math.Round(float64(I) * EF))
		}
		n++
		randVal := 2.0 * rand.Float64()
		EF = EF + (0.1 - randVal*(0.08+randVal*0.02))
	} else {
		n = 0
		I = 1
		oldI = 0
		randVal := 3.0 + 2.0*rand.Float64()
		EF = EF + (0.1 - randVal*(0.08+randVal*0.02))
	}
	if EF < 1.3 {
		EF = 1.3
	}
	card.LastReview = timeutil.TruncateToDay(time.Now())
	card.NextReview = card.LastReview.AddDate(0, 0, oldI)
	card.Sm2N = n
	card.Sm2EF = EF
	card.Sm2I = I
	card.NumReviews++
	return card
}
