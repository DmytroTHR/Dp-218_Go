package services

import (
	"Dp218Go/models"
	"Dp218Go/repositories"
)

type ScooterService struct {
	repoScooter repositories.ScooterRepo
}


func NewScooterService(repoScooter repositories.ScooterRepo) *ScooterService {
	return &ScooterService{repoScooter: repoScooter}
}

func (ser *ScooterService) GetAllScooters() (*models.ScooterListDTO, error) {
	return ser.repoScooter.GetAllScooters()
}

func (ser *ScooterService) GetScooterById(uid int) (models.ScooterDTO, error) {
	return ser.repoScooter.GetScooterById(uid)
}

func (ser *ScooterService) GetScooterStatus(scooterID int) (models.ScooterStatus, error) {
	return ser.repoScooter.GetScooterStatus(scooterID)
}

func (ser *ScooterService) SendCurrentPosition(id int, lat, lon float64) error {
	return ser.repoScooter.SendCurrentPosition(id, lat, lon)
}

