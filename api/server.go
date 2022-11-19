package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteOptions struct {
	Cfg *gorm.DB
}

func New(opt *RouteOptions) *gin.Engine {
	router := gin.Default()

	h := RouteOptions{
		Cfg: opt.Cfg,
	}

	apiv1 := router.Group("/users")
	{
		apiv1.POST("/sign-up", h.SignUp)
		apiv1.POST("/sign-in", h.SignIn)
		apiv1.GET("/validate", h.ReQuireAuth, h.Validate)
	}

	return router
}