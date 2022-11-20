package services

import (
	"context"
	"errors"

	"restapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarServiceImpl struct {
	carCollection *mongo.Collection
	ctx context.Context
}

func NewCarService(carCollection *mongo.Collection, ctx context.Context) CarService {
	return &CarServiceImpl {
		carCollection: carCollection,
		ctx: ctx,
	}
}

func (c *CarServiceImpl) GetAll() ([]*models.Car, error) {
	var cars []*models.Car
	cursor, err := c.carCollection.Find(c.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var car models.Car
		err := cursor.Decode(&car)
		if err != nil {
			return nil, err
		}
		cars = append(cars, &car)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(c.ctx)
	if len(cars) == 0 {
		return nil, errors.New("No documents found!")
	}
	return cars, nil
}

func (c *CarServiceImpl) Add(car *models.Car) error {
	_, err := c.carCollection.InsertOne(c.ctx, &models.Car{
		Id: primitive.NewObjectID(),
		Name: car.Name,
	})
	return err
}

func (c *CarServiceImpl) Update(car *models.Car, id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "id", Value: objID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: car.Name}}}}
	result, _ := c.carCollection.UpdateOne(c.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update!")
	}
	return nil
}

func (c *CarServiceImpl) Delete(id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "id", Value: objID}}
	result, _ := c.carCollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched document found for delete!")
	}
	return nil
}