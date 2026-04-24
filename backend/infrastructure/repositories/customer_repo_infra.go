package repositories

import (
	"bank-app/domain/entities"

	"gorm.io/gorm"
)

type CustomerRepoInfra struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepoInfra {
	return &CustomerRepoInfra{db}
}

func (r *CustomerRepoInfra) Create(c *entities.Customer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(c).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *CustomerRepoInfra) FindAll() ([]entities.Customer, error) {
	var list []entities.Customer
	return list, r.db.Find(&list).Error
}

func (r *CustomerRepoInfra) FindByID(publicID string) (*entities.Customer, error) {
	var customer entities.Customer

	return &customer, r.db.Where("public_id = ?", publicID).First(&customer).Error
}
