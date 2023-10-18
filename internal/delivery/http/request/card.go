package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCardRequest struct {
	UserID       primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	DeckID       primitive.ObjectID `json:"deck_id" binding:"required"`
	Question     string             `json:"question" binding:"required"`
	Answer       string             `json:"answer" binding:"required"`
	WrongAnswers []string           `json:"wrong_answers" binding:"required"`
}
