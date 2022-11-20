package services

import "restapi/models"

type CarService interface {
	GetAll() ([]*models.Car, error)
	Add(*models.Car) error
	Update(*models.Car, *string) error
	Delete(*string) error
}