package repositories

import (
	"bank-app/domain/entities"

	"gorm.io/gorm"
)

type AccountRepoInfra struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepoInfra {
	return &AccountRepoInfra{db}
}

func (r *AccountRepoInfra) Create(a *entities.Account) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(a).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *AccountRepoInfra) FindByID(publicID string) (*entities.Account, error) {
	var acc entities.Account

	err := r.db.Where("public_id = ?", publicID).Find(&acc).Error
	//err := r.db.Find(&acc, "publicID = ?", publicID)

	return &acc, err

}

func (r *AccountRepoInfra) Update(a *entities.Account) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(a).Error; err != nil {
			return err
		}

		return nil
	})
}
