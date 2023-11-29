package entity

import (
	"time"
	"vietcard-backend/pkg/timeutil"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deck struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt           time.Time          `json:"created_at" bson:"created_at"`
	UserID              primitive.ObjectID `json:"user_id" bson:"user_id"`
	IsPublic            bool               `json:"is_public" bson:"is_public"`
	IsFavorite          bool               `json:"is_favorite" bson:"is_favorite"`
	Name                string             `json:"name" bson:"name"`
	Description         string             `json:"description" bson:"description"`
	DescriptionImageURL string             `json:"description_img_url" bson:"description_img_url"`
	Views               int                `json:"views" bson:"views"`
	Rating              float32            `json:"rating" bson:"rating"`
	TotalCards          int                `json:"total_cards" bson:"total_cards"`
	TotalLearnedCards   int                `json:"total_learned_cards" bson:"total_learned_cards"`
	MaxNewCards         int                `json:"max_new_cards" bson:"max_new_cards"`
	MaxReviewCards      int                `json:"max_review_cards" bson:"max_review_cards"`
	LastReview          time.Time          `json:"last_review" bson:"last_review"`
	CurNewCards         int                `json:"cur_new_cards" bson:"cur_new_cards"`
	CurReviewCards      int                `json:"cur_review_cards" bson:"cur_review_cards"`
}

type DeckWithReviewCards struct {
	Deck  `bson:"inline"`
	Cards *[]Card `json:"cards" bson:"cards"`
}

func (deck *Deck) SetDefault() *Deck {
	deck.CreatedAt = time.Now()
	deck.MaxNewCards = 20
	deck.MaxReviewCards = 100
    deck.Rating = 5
    deck.Views = 1
    deck.TotalLearnedCards = 0
    deck.IsFavorite = false
	return deck
}

func (deck *Deck) UpdateReview() *Deck {
	cur := timeutil.TruncateToDay(time.Now())
	if cur.Equal(deck.LastReview) {
		return deck
	}
	deck.LastReview = cur
	deck.CurNewCards = 0
	deck.CurReviewCards = 0
	return deck
}
