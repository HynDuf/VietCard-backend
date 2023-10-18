package deckrepo

import (
	"context"
	"errors"
	"vietcard-backend/internal/domain/entity"
	"vietcard-backend/internal/domain/interface/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type deckRepository struct {
	db      *mongo.Database
	colName string
}

func NewDeckRepository(db *mongo.Database) repository.DeckRepository {
	return &deckRepository{
		db:      db,
		colName: "decks",
	}
}

func (dr *deckRepository) CreateDeck(deck *entity.Deck) error {
	deck.SetDefault()
	_, err := dr.db.Collection(dr.colName).InsertOne(context.TODO(), deck)
	if err != nil {
		return err
	}
	return nil
}

func (dr *deckRepository) GetDeckByID(id *string) (*entity.Deck, error) {
	oID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}
	var deck entity.Deck
	err = dr.db.Collection(dr.colName).FindOne(context.TODO(), bson.D{{Key: "_id", Value: oID}}).Decode(&deck)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &deck, nil
}

func (dr *deckRepository) UpdateDeck(deck *entity.Deck) error {
	oID := deck.ID
	var newDeck entity.Deck
	err := dr.db.Collection(dr.colName).FindOneAndReplace(context.TODO(), bson.D{{Key: "_id", Value: oID}}, deck).Decode(&newDeck)
	if err != nil {
		return err
	}
	if newDeck.ID != deck.ID {
		err = errors.New("error: updated card had different ID")
		return err
	}
	return nil
}

func (dr *deckRepository) GetCardsAllDecksOfUser(userID *string) (*[]entity.DeckWithReviewCards, error) {
	// Requires the MongoDB Go Driver
	// https://go.mongodb.org/mongo-driver
	ctx := context.TODO()

	uID, err := primitive.ObjectIDFromHex(*userID)
	if err != nil {
		return nil, err
	}

	// Open an aggregation cursor
	coll := dr.db.Collection(dr.colName)
	cursor, err := coll.Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"user_id", uID}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "cards"},
					{"localField", "_id"},
					{"foreignField", "deck_id"},
					{"as", "cards"},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var deckWithCards []entity.DeckWithReviewCards
	if err = cursor.All(context.TODO(), &deckWithCards); err != nil {
		return nil, err
	}
	return &deckWithCards, nil
}
