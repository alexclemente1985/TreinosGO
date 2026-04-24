package customer

import (
	handlers "bank-app/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, cH *handlers.CustomerHandler) *gin.Engine {

	r.POST("/customers", cH.Create)
	r.GET("/customers", cH.List)

	return r
}

func RegisterRouterGroup(mainGroup *gin.RouterGroup, cH *handlers.CustomerHandler) {
	group := mainGroup.Group("customers")
	{
		group.POST("/", cH.Create)
		group.GET("/", cH.List)
	}
}
