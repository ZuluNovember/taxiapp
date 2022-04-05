package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Driver struct {
	ID       string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Location Geometry `json:"location" bson:"location"`
}

type Geometry struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

var maxDistance = 5000

func createDriver(lon, lat float64) Driver {
	d := Driver{Location: Geometry{
		Type:        "Point",
		Coordinates: []float64{lon, lat},
	},
	}
	return d
}

func AddDriver(lon, lat float64) (*mongo.InsertOneResult, error) {
	ctx, c := context.WithCancel(context.Background())
	defer c()

	d := createDriver(lon, lat)

	res, err := locationCollection.InsertOne(ctx, d)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func BulkAddDriver(coordinates [][]float64) (*mongo.BulkWriteResult, error) {
	ctx, c := context.WithCancel(context.Background())
	defer c()

	docs := make([]mongo.WriteModel, len(coordinates))

	for i, v := range coordinates {
		d := createDriver(v[0], v[1])
		docs[i] = mongo.NewInsertOneModel().SetDocument(d)
	}

	res, err := locationCollection.BulkWrite(ctx, docs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindDrivers(lon, lat float64) ([]Driver, error) {
	ctx, c := context.WithCancel(context.Background())
	defer c()

	var closeDrivers []Driver

	geometry := Geometry{
		Type:        "Point",
		Coordinates: []float64{lon, lat},
	}

	filter := bson.D{
		{Key: "location",
			Value: bson.D{
				{Key: "$near", Value: bson.D{
					{Key: "$geometry", Value: geometry},
					{Key: "$maxDistance", Value: maxDistance},
				}},
			}},
	}

	cursor, err := locationCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var d Driver

		err := cursor.Decode(&d)
		if err != nil {
			return nil, err
		}
		closeDrivers = append(closeDrivers, d)
	}

	return closeDrivers, nil
}
