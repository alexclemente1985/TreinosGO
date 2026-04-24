package mappers

import (
	"bank-app/delivery/http/dtos"
	"bank-app/domain/entities"
)

func ToUserEntity(d dtos.UserDTO) *entities.User {
	u := &entities.User{
		PublicID: d.ID,
		Email:    d.Email,
		Password: d.Password,
	}

	return u
}

// func ToUserDTO(u *entities.User) dtos.UserDTO {
// 	res := dtos.UserDTO{
// 		ID:          u.PublicID,
// 		Email:       u.Email,
// 		MobilePhone: u.MobilePhone,
// 		CreatedAt:   u.CreatedAt.String(),
// 	}
// 	return res
// }

// func ToAuthDTO(a *entities.Auth) dtos.AuthDTO {
// 	res := dtos.AuthDTO{
// 		AccessToken: a.accessToken,
// 		Name:        a.Name,
// 		Email:       a.Email,
// 		Phone:       a.Phone,
// 		MobilePhone: a.MobilePhone,
// 		CreatedAt:   a.CreatedAt.String(),
// 	}
// 	return res
// }
