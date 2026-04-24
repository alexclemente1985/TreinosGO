package usecase

import (
	"bank-app/domain/entities"
	"bank-app/domain/repositories"
)

type CustomerUseCase struct {
	repo repositories.CustomerRepository
}

func NewCustomerUseCase(r repositories.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{r}
}

func (u *CustomerUseCase) Create(c *entities.Customer) error {
	return u.repo.Create(c)
}

func (u *CustomerUseCase) List() ([]entities.Customer, error) {
	return u.repo.FindAll()
}

func (u *CustomerUseCase) FindByID(publicID string) (*entities.Customer, error) {
	customer, error := u.repo.FindByID(publicID)
	return customer, error
}
