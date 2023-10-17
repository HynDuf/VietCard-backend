package userrepo

import (
	"context"
	"time"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db      *mongo.Database
	colName string
}

func NewUserRepository(db *mongo.Database) repository.UserRepository {
	return &userRepository{
		db:      db,
		colName: "users",
	}
}

func (ur *userRepository) Create(user *entity.User) error {
	timeNow := time.Now()
	user.CreatedAt = timeNow
	_, err := ur.db.Collection(ur.colName).InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetByEmail(email *string) (*entity.User, error) {
	var user entity.User
	err := ur.db.Collection(ur.colName).FindOne(context.TODO(), bson.D{{Key: "email", Value: *email}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetByID(id *string) (*entity.User, error) {
	var user entity.User
	err := ur.db.Collection(ur.colName).FindOne(context.TODO(), bson.D{{Key: "_id", Value: *id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
