package mongodb

import (
	"context"
	"log"
	"vietcard-backend/bootstrap"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBConnection(URI string) *mongo.Database {
	// Set client options
	clientOptions := options.Client().ApplyURI(URI)

	// Connect to MongoDB
	var err error
    client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
    return client.Database(bootstrap.E.MongoDBName)
}
