package services

import (
	"Dp218Go/models"
	"Dp218Go/repositories"
)

type OrderService struct {
	repoOrder   repositories.OrderRepo
	scooterRepo repositories.ScooterRepo
}

func NewOrderService(orderRepo repositories.OrderRepo, scooterRepo repositories.ScooterRepo) *OrderService {
	return &OrderService{repoOrder: orderRepo, scooterRepo: scooterRepo}
}

func (os *OrderService) CreateOrder(user models.User, scooterID, startID, endID int, distance float64) (models.Order, error) {
	return os.repoOrder.CreateOrder(user, scooterID, startID, endID, distance)
}

func (os *OrderService) GetAllOrders() (*models.OrderList, error) {
	return os.repoOrder.GetAllOrders()
}
//TODO continue implementation
