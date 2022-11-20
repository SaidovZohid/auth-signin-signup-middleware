package api

import (
	"github.com/SaidovZohid/auth-signin-signup-middleware/config"
	"github.com/SaidovZohid/auth-signin-signup-middleware/storage"
	"github.com/gin-gonic/gin"
)

type RouteOptions struct {
	Cfg *config.Config
	Storage storage.StorageI
}

// @title           Swagger for blog api
// @version         2.0
// @description     This is a blog service api.
// @host      localhost:8080
// @BasePath  /users
func New(opt *RouteOptions) *gin.Engine {
	router := gin.Default()

	h := RouteOptions{
		Cfg: opt.Cfg,
		Storage: opt.Storage,
	}

	apiv1 := router.Group("/users")
	{
		apiv1.POST("/sign-up", h.SignUp)
		apiv1.POST("/sign-in", h.SignIn)
		apiv1.GET("/validate", h.ReQuireAuth, h.Validate)
	}

	return router
}