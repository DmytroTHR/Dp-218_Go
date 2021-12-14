package repositories

import "Dp218Go/models"

type OrderRepo interface {
	CreateOrder(user models.User, scooterID, startID, endID int, distance float64) (models.Order, error)
	GetAllOrders() (*models.OrderList, error)
	//SetOrderStart(order *models.Order, status models.ScooterStatusInRent) error
	//SetOrderEnd(order *models.Order, status models.ScooterStatusInRent) error
	//UpdateOrder(order *models.Order) error
}

//TODO implement other methods