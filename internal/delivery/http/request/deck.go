package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateDeckRequest struct {
	UserID primitive.ObjectID `json:"user_id" swaggerignore:"true"`
	Name   string             `json:"name" binding:"required"`
}

type UpdateDeckRequest struct {
	DeckID   *primitive.ObjectID `json:"deck_id" bson:"_id,omitempty" binding:"required"`
	IsGlobal *bool               `json:"is_global" bson:"is_global,omitempty"`
	Name     *string             `json:"name" bson:"name,omitempty"`
}
