package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Fact struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content string             `json:"content" bson:"content"`
}
