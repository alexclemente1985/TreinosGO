package handlers

import (
	"bank-app/delivery/http/dtos"
	"bank-app/delivery/http/mappers"
	"bank-app/usecase"
	"fmt"
	"net/http"

	ulid "bank-app/shared/utils/id"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	uc *usecase.CustomerUseCase
}

func NewCustomerHandler(u *usecase.CustomerUseCase) *CustomerHandler {
	return &CustomerHandler{u}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var input dtos.CustomerDTO
	c.ShouldBindJSON(&input)
	input.ID = ulid.New()

	entity := mappers.ToCustomerEntity(input)

	h.uc.Create(entity)

	fmt.Printf("ENTITY - %s", entity)

	c.JSON(http.StatusCreated, mappers.ToCustomerDTO(entity))
}

func (h *CustomerHandler) List(c *gin.Context) {
	list, error := h.uc.List()

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	var dtoList []dtos.CustomerDTO

	for _, c := range list {
		dtoList = append(dtoList, mappers.ToCustomerDTO(&c))
	}

	c.JSON(http.StatusOK, dtoList)

}
