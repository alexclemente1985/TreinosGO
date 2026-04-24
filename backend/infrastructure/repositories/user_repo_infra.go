package repositories

import (
	"bank-app/domain/entities"

	"gorm.io/gorm"
)

type UserRepoInfra struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoInfra {
	return &UserRepoInfra{db}
}

func (r *UserRepoInfra) Create(u *entities.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(u).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *UserRepoInfra) FindByID(publicID string) (*entities.User, error) {
	var user entities.User

	return &user, r.db.Where("public_id = ?", publicID).First(&user).Error
}

func (r *UserRepoInfra) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	return &user, r.db.Where("email = ?", email).First(&user).Error
}
