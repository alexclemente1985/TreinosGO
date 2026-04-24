package auth

import (
	"bank-app/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, authH *handlers.AuthHandler) *gin.Engine {

	r.POST("/register", authH.Register)
	r.POST("/login", authH.Login)

	return r
}

func RegisterRouterGroup(mainGroup *gin.RouterGroup, authH *handlers.AuthHandler) {
	group := mainGroup.Group("auth")
	{
		group.POST("/register", authH.Register)
		group.POST("/login", authH.Login)
	}
}
