package request

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateDeckRequest struct {
	UserID      primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description"`
}

type UpdateDeckRequest struct {
	DeckID         *primitive.ObjectID `json:"deck_id" bson:"_id,omitempty" binding:"required"`
	IsGlobal       *bool               `json:"is_global" bson:"is_global,omitempty"`
	Name           *string             `json:"name" bson:"name,omitempty"`
	MaxNewCards    *int                `json:"max_new_cards" bson:"max_new_cards,omitempty"`
	MaxReviewCards *int                `json:"max_review_cards" bson:"max_review_cards,omitempty"`
	LastReview     *time.Time          `json:"last_review" bson:"last_review,omitempty"`
	CurNewCards    *int                `json:"cur_new_cards" bson:"cur_new_cards,omitempty"`
	CurReviewCards *int                `json:"cur_review_cards" bson:"cur_review_cards,omitempty"`
}

type CopyDeckRequest struct {
	DeckID *primitive.ObjectID `json:"deck_id" binding:"required"`
}
