package cardrepo

import (
	"context"
	"errors"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type cardRepository struct {
	db      *mongo.Database
	colName string
}

func NewCardRepository(db *mongo.Database) repository.CardRepository {
	return &cardRepository{
		db:      db,
		colName: "cards",
	}
}

func (cr *cardRepository) CreateCard(card *entity.Card) error {
	card.SetDefault()
	_, err := cr.db.Collection(cr.colName).InsertOne(context.TODO(), card)
	if err != nil {
		return err
	}
	return nil
}

func (cr *cardRepository) GetCardByID(id *string) (*entity.Card, error) {
	oID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}
	var card entity.Card
	err = cr.db.Collection(cr.colName).FindOne(context.TODO(), bson.D{{Key: "_id", Value: oID}}).Decode(&card)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &card, nil
}

func (cr *cardRepository) UpdateCard(card *entity.Card) error {
	oID := card.ID
	var newCard entity.Card
	err := cr.db.Collection(cr.colName).FindOneAndReplace(context.TODO(), bson.D{{Key: "_id", Value: oID}}, card).Decode(&newCard)
	if err != nil {
		return err
	}
	if newCard.ID != card.ID {
		err = errors.New("error: updated card had different ID")
		return err
	}
	return nil
}

func (cr *cardRepository) UpdateCardReview(card *entity.Card) error {
	oID := card.ID
	filter := bson.D{{Key: "_id", Value: oID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "sm2_n", Value: card.Sm2N},
			{Key: "sm2_ef", Value: card.Sm2EF},
			{Key: "sm2_i", Value: card.Sm2I},
		}}}
	_, err := cr.db.Collection(cr.colName).UpdateOne(context.TODO(), &filter, &update)
	if err != nil {
		return err
	}
	return nil
}
