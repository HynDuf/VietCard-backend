package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deck struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	IsGlobal  bool               `json:"is_global" bson:"is_global"`
	Name      string             `json:"name" bson:"name"`
}

type DeckWithReviewCards struct {
	Deck  `bson:"inline"`
	Cards []Card `json:"cards" bson:"cards"`
}

func (deck *Deck) SetDefault() *Deck {
	deck.CreatedAt = time.Now()
	return deck
}
