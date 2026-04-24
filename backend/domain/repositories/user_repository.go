package repositories

import "bank-app/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
}
