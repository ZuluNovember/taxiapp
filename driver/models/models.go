package models

import (
	"context"
	"log"
	"time"

	"github.com/ZuluNovember/taxiapp/driver/pkg/settings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var locationCollection *mongo.Collection

func Setup() {
	ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
	defer c()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(settings.Configuration.Database.Address))
	if err != nil {
		log.Fatal(err)
	}

	locationCollection = client.Database(settings.Configuration.Database.Name).Collection("locations")
	err = createIndexes()
	if err != nil {
		log.Fatal(err)
	}
}

func CloseConnection() {
	ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
	defer c()

	client.Disconnect(ctx)
}

func createIndexes() error {
	ctx, c := context.WithCancel(context.Background())
	defer c()
	locationIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "location", Value: "2dsphere"}},
	}

	_, err := locationCollection.Indexes().CreateOne(ctx, locationIndex)
	if err != nil {
		return err
	}
	return nil
}
