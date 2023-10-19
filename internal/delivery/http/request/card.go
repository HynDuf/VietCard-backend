package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCardRequest struct {
	UserID       primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	DeckID       primitive.ObjectID `json:"deck_id" binding:"required"`
	Question     string             `json:"question" binding:"required"`
	Answer       string             `json:"answer" binding:"required"`
	WrongAnswers []string           `json:"wrong_answers" binding:"required"`
}

type UpdateCardRequest struct {
	CardID       *primitive.ObjectID `json:"card_id" bson:"_id,omitempty" binding:"required"`
	DeckID       *primitive.ObjectID `json:"deck_id" bson:"deck_id,omitempty"`
	Question     *string             `json:"question" bson:"question,omitempty"`
	Answer       *string             `json:"answer" bson:"answer,omitempty"`
	WrongAnswers *[]string           `json:"wrong_answers" bson:"wrong_answers,omitempty"`
}

type UpdateReviewCardsRequest struct {
	DeckID    primitive.ObjectID   `json:"deck_id" binding:"required"`
	TotalXP   int                  `json:"total_xp" binding:"required"`
	CardIDs   []primitive.ObjectID `json:"card_ids" binding:"required"`
	IsCorrect []bool               `json:"is_correct" binding:"required"`
}
