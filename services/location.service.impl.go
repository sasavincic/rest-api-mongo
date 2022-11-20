package services

import (
	"context"
	"errors"

	"restapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationServiceImpl struct {
	locationCollection *mongo.Collection
	ctx context.Context
}

func NewLocationService(locationCollection *mongo.Collection, ctx context.Context) LocationService {
	return &LocationServiceImpl {
		locationCollection: locationCollection,
		ctx: ctx,
	}
}

func (l *LocationServiceImpl) GetAll() ([]*models.Location, error) {
	var locations []*models.Location
	cursor, err := l.locationCollection.Find(l.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(l.ctx) {
		var location models.Location
		err := cursor.Decode(&location)
		if err != nil {
			return nil, err
		}
		locations = append(locations, &location)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(l.ctx)
	if len(locations) == 0 {
		return nil, errors.New("No documents found!")
	}
	return locations, nil
}

func (l *LocationServiceImpl) Add(location *models.Location) error {
	_, err := l.locationCollection.InsertOne(l.ctx, &models.Location{
		Id: primitive.NewObjectID(),
		Name: location.Name,
	})
	return err
}

func (l *LocationServiceImpl) Update(location *models.Location, id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "id", Value: objID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: location.Name}}}}
	result, _ := l.locationCollection.UpdateOne(l.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update!")
	}
	return nil
}

func (l *LocationServiceImpl) Delete(id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "id", Value: objID}}
	result, _ := l.locationCollection.DeleteOne(l.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched document found for delete!")
	}
	return nil
}