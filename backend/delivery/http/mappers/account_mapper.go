package mappers

import (
	"bank-app/delivery/http/dtos"
	"bank-app/domain/entities"
)

func ToAccountEntity(d dtos.AccountDTO, customerID int64) *entities.Account {
	a := &entities.Account{
		PublicID:   d.ID,
		CustomerID: customerID,
		Type:       d.Type,
		Balance:    d.Balance,
	}

	return a
}

func ToAccountDTO(a *entities.Account, customerPublicID string) dtos.AccountDTO {
	res := dtos.AccountDTO{
		ID:         a.PublicID,
		Type:       a.Type,
		Balance:    a.Balance,
		CustomerID: customerPublicID,
	}

	return res
}
