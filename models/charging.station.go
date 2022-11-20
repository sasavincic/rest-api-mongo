package models


import "go.mongodb.org/mongo-driver/bson/primitive"

type ChargingStation struct {
	Id primitive.ObjectID `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}