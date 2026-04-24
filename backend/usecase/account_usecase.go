package usecase

import (
	"bank-app/domain/entities"
	"bank-app/domain/repositories"
	"errors"
	"fmt"
)

type AccountUseCase struct {
	repo repositories.AccountRepository
}

func NewAccountUseCase(r repositories.AccountRepository) *AccountUseCase {
	return &AccountUseCase{r}
}

func (u *AccountUseCase) Create(a *entities.Account) error {
	return u.repo.Create(&entities.Account{
		CustomerID: a.CustomerID,
		Type:       a.Type,
		PublicID:   a.PublicID,
		Balance:    0,
	})
}

func (u *AccountUseCase) Deposit(publicID string, amount float64) error {
	acc, _ := u.repo.FindByID(publicID)
	acc.Balance += amount
	return u.repo.Update(acc)
}

func (u *AccountUseCase) Withdraw(publicID string, amount float64) error {
	acc, _ := u.repo.FindByID(publicID)

	if acc.Balance < amount {
		fmt.Println("Falta de grana...", acc.Balance, "----", amount)
		return errors.New("insufficient funds")
	}

	acc.Balance -= amount
	return u.repo.Update(acc)
}

func (u *AccountUseCase) Balance(publicID string) (float64, error) {
	acc, _ := u.repo.FindByID(publicID)
	return acc.Balance, nil
}
