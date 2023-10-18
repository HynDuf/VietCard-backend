package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateDeckRequest struct {
	UserID primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	Name   string             `json:"name" binding:"required"`
}

