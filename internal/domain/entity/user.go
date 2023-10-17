package entity

import "time"

type User struct {
	ID             int       `json:"id" bson:"_id"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
	Name           string    `json:"name" bson:"name"`
	Email          string    `json:"email" bson:"email"`
	HashedPassword string    `json:"hashed_password" bson:"hashed_password"`
}
