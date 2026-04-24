package usecase

import (
	"bank-app/delivery/http/dtos"
	"bank-app/domain/entities"
	"bank-app/domain/repositories"
	"errors"
	"fmt"
	"time"

	authConfig "bank-app/config/auth"
	ulid "bank-app/shared/utils/id"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	repo repositories.UserRepository
}

func NewAuthUseCase(r repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{r}
}

func (u *AuthUseCase) Register(user *entities.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	return u.repo.Create(&entities.User{
		PublicID: ulid.New(),
		Email:    user.Email,
		Password: string(hash),
	})
}

func (u *AuthUseCase) Login(user *entities.User) (*dtos.AuthDTO, error) {
	registeredUser, err := u.repo.FindByEmail(user.Email)

	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(registeredUser.Password), []byte(user.Password))

	if err != nil {
		return nil, fmt.Errorf("Invalid Credentials: %s", err)
	}

	authDTO := dtos.AuthDTO{
		AccessToken:  generateToken(user.PublicID, 15*time.Minute),
		RefreshToken: generateToken(user.PublicID, 7*24*time.Hour),
	}

	return &authDTO, nil
}

func generateToken(userID string, duration time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(duration).Unix(),
	})

	t, _ := token.SignedString(authConfig.JwtSecret)

	return t
}
