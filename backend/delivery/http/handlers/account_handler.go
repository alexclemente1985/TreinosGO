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

type AccountHandler struct {
	ucAccount  *usecase.AccountUseCase
	ucCustomer *usecase.CustomerUseCase
}

func NewAccountHandler(uca *usecase.AccountUseCase, ucc *usecase.CustomerUseCase) *AccountHandler {
	return &AccountHandler{uca, ucc}
}

func (h *AccountHandler) Create(c *gin.Context) {
	var input dtos.AccountDTO
	c.ShouldBindJSON(&input)
	input.ID = ulid.New()

	customer, error := h.ucCustomer.FindByID(input.CustomerID)

	if error != nil {
		c.JSON(http.StatusNotFound, "No customers with this publicID.")
		return
	}

	entity := mappers.ToAccountEntity(input, customer.ID)

	if error := h.ucAccount.Create(entity); error != nil {
		c.JSON(http.StatusInternalServerError, error)
		return
	}

	c.JSON(http.StatusCreated, entity)

}

func (h *AccountHandler) Deposit(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found or invalid")
		return
	}

	var input dtos.AccountTransactionDTO
	if error := c.ShouldBindJSON(&input); error != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid JSON structure: %s", error))
		return
	}

	var err = h.ucAccount.Deposit(id, input.Amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Transaction completed with success")

}

func (h *AccountHandler) Withdraw(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found or invalid")
		return
	}

	var input dtos.AccountTransactionDTO
	if error := c.ShouldBindJSON(&input); error != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON structure")
		return
	}

	var err = h.ucAccount.Withdraw(id, input.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Withdraw transaction completed with success")

}

func (h *AccountHandler) Balance(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found or invalid")
		return
	}

	balance, err := h.ucAccount.Balance(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var balanceDTO = dtos.AccountBalanceDTO{
		AccountID: id,
		Balance:   balance,
	}

	c.JSON(http.StatusOK, gin.H{"balanceDTO": balanceDTO})

}
