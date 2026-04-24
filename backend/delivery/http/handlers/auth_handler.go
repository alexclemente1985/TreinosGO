package handlers

import (
	"bank-app/delivery/http/dtos"
	"bank-app/delivery/http/mappers"
	ulid "bank-app/shared/utils/id"
	"bank-app/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc *usecase.AuthUseCase
}

func NewAuthHandler(u *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{u}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input dtos.UserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid login data: %s", err.Error())})
		return
	}

	input.ID = ulid.New()

	entity := mappers.ToUserEntity(input)

	if err := h.uc.Register(entity); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error: %s", err))
		return
	}

	c.JSON(http.StatusCreated, "User successfully created.")

}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dtos.UserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid login data")
		return
	}

	entity := mappers.ToUserEntity(input)

	authDTO, err := h.uc.Login(entity)

	if err != nil {
		c.JSON(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
		return
	}

	c.JSON(http.StatusOK, authDTO)

}
