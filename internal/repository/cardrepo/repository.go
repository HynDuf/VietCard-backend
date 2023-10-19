package cardrepo

import (
	"context"
	"vietcard-backend/internal/delivery/http/request"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (cr *cardRepository) UpdateCard(cardID *string, req *request.UpdateCardRequest) (*entity.Card, error) {
	cID, err := primitive.ObjectIDFromHex(*cardID)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: cID}}
	update := bson.D{{Key: "$set", Value: *req}}
	option := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedCard entity.Card
	err = cr.db.Collection(cr.colName).FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&updatedCard)
	if err != nil {
		return nil, err
	}
	return &updatedCard, nil
}

func (cr *cardRepository) GetCardsByDeck(deckID *string) (*[]entity.Card, error) {
	dID, err := primitive.ObjectIDFromHex(*deckID)
	if err != nil {
		return nil, err
	}
	cursor, err := cr.db.Collection(cr.colName).Find(
		context.TODO(),
		bson.D{{Key: "deck_id", Value: dID}},
	)
	if err != nil {
		return nil, err
	}

	var cards []entity.Card
	if err = cursor.All(context.TODO(), &cards); err != nil {
		return nil, err
	}

	return &cards, nil
}
