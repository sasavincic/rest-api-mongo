package services

import (
	"context"
	"errors"

	"restapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChargingStationServiceImpl struct {
	chargingStationCollection *mongo.Collection
	ctx context.Context
}

func NewChargingStationService(chargingStationCollection *mongo.Collection, ctx context.Context) ChargingStationService {
	return &ChargingStationServiceImpl {
		chargingStationCollection: chargingStationCollection,
		ctx: ctx,
	}
}

func (c *ChargingStationServiceImpl) GetAll() ([]*models.ChargingStation, error) {
	var chargingStations []*models.ChargingStation
	cursor, err := c.chargingStationCollection.Find(c.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var chargingStation models.ChargingStation
		err := cursor.Decode(&chargingStation)
		if err != nil {
			return nil, err
		}
		chargingStations = append(chargingStations, &chargingStation)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(c.ctx)
	if len(chargingStations) == 0 {
		return nil, errors.New("No documents found!")
	}
	return chargingStations, nil
}

func (c *ChargingStationServiceImpl) Add(chargingStation *models.ChargingStation) error {
	_, err := c.chargingStationCollection.InsertOne(c.ctx, &models.ChargingStation{
		Id: primitive.NewObjectID(),
		Name: chargingStation.Name,
	})
	return err
}

func (c *ChargingStationServiceImpl) Update(chargingStation *models.ChargingStation, id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "id", Value: objID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: chargingStation.Name}}}}
	result, _ := c.chargingStationCollection.UpdateOne(c.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update!")
	}
	return nil
}

func (c *ChargingStationServiceImpl) Delete(id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "id", Value: objID}}
	result, _ := c.chargingStationCollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched document found for delete!")
	}
	return nil
}