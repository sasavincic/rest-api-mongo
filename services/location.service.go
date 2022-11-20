package services

import "restapi/models"

type LocationService interface {
	GetAll() ([]*models.Location, error)
	Add(*models.Location) error
	Update(*models.Location, *string) error
	Delete(*string) error
}