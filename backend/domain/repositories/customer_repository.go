package repositories

import "bank-app/domain/entities"

type CustomerRepository interface {
	Create(customer *entities.Customer) error
	FindAll() ([]entities.Customer, error)
	FindByID(publicID string) (*entities.Customer, error)
}
