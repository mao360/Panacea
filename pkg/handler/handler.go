package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) InitRouts() *echo.Echo {
	router := echo.New()
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		api.GET("/home", nil)
		api.GET("/selection", nil)
		api.GET("/forDoctors", nil)
		api.GET("/search", nil)
	}
	return router
}
