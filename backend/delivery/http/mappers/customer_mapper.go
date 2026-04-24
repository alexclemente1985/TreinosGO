package mappers

import (
	"bank-app/delivery/http/dtos"
	"bank-app/domain/entities"
)

func ToCustomerEntity(d dtos.CustomerDTO) *entities.Customer {
	c := &entities.Customer{
		PublicID:    d.ID,
		Name:        d.Name,
		Email:       d.Email,
		Phone:       d.Phone,
		MobilePhone: d.MobilePhone,
	}

	return c
}

func ToCustomerDTO(c *entities.Customer) dtos.CustomerDTO {
	res := dtos.CustomerDTO{
		ID:          c.PublicID,
		Name:        c.Name,
		Email:       c.Email,
		Phone:       c.Phone,
		MobilePhone: c.MobilePhone,
		CreatedAt:   c.CreatedAt.String(),
	}
	return res
}
