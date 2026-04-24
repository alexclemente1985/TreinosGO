package repositories

import "bank-app/domain/entities"

type AccountRepository interface {
	Create(account *entities.Account) error
	FindByID(id string) (*entities.Account, error)
	Update(account *entities.Account) error
}
