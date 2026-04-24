package routes

import (
	handlers "bank-app/delivery/http/handlers"
	"bank-app/delivery/middleware"

	accountRouter "bank-app/delivery/routes/account"
	authRouter "bank-app/delivery/routes/auth"
	customerRouter "bank-app/delivery/routes/customer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cH *handlers.CustomerHandler, aH *handlers.AccountHandler, authH *handlers.AuthHandler) *gin.Engine {

	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false
	r.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}))

	v1 := r.Group("/api/v1")

	authRouter.RegisterRouterGroup(v1, authH)

	v1.Use(middleware.AuthMiddleware())
	{
		customerRouter.RegisterRouterGroup(v1, cH)
		accountRouter.RegisterRouterGroup(v1, aH)
	}

	// authRouter.SetupRouter(r, authH)
	// customerRouter.SetupRouter(r, cH)
	// accountRouter.SetupRouter(r, aH)

	return r
}
