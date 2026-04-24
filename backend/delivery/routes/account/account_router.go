package account

import (
	handlers "bank-app/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, aH *handlers.AccountHandler) *gin.Engine {

	r.POST("/accounts", aH.Create)
	r.POST("/accounts/:id/deposit", aH.Deposit)
	r.POST("/accounts/:id/withdraw", aH.Withdraw)
	r.GET("/accounts/:id/balance", aH.Balance)

	return r
}

func RegisterRouterGroup(mainGroup *gin.RouterGroup, aH *handlers.AccountHandler) {
	group := mainGroup.Group("accounts")
	{
		group.POST("/", aH.Create)
		group.POST("/:id/deposit", aH.Deposit)
		group.POST("/:id/withdraw", aH.Withdraw)
		group.GET("/:id/balance", aH.Balance)
	}
}
