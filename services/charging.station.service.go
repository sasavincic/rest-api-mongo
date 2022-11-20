package services

import "restapi/models"

type ChargingStationService interface {
	GetAll() ([]*models.ChargingStation, error)
	Add(*models.ChargingStation) error
	Update(*models.ChargingStation, *string) error
	Delete(*string) error
}