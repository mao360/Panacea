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
		//ОБЩИЕ ОБРАБОТЧИКИ
		api.DELETE("/home/profile", h.deleteProfile)
		api.GET("/articles", h.readPost)
		api.GET("/:id", h.checkProfile)
		api.POST("/discussions", h.createComment)
		api.PUT("/discussions", h.editComment)
		api.DELETE("/discussions", h.deleteComment)

		//ВРАЧА ОБРАБОТЧИКИ
		api.POST("/!!!", h.acceptPatient)
		api.DELETE("/!!!", h.unAcceptPatient)
		api.POST("/!!!", h.ratePatient)
		api.POST("/posts", h.createPost)
		api.PUT("/posts", h.editPost)
		api.DELETE("/posts", h.deletePost)

		//ПАЦИЕНТА ОБРАБОТЧИКИ
		api.POST("/!!!", h.createRequest)
		api.DELETE("/!!!", h.deleteRequest)
		api.POST("/!!!", h.rateDoctor)
		api.POST("/discussions", h.createDiscussion)
		api.PUT("/discussions", h.editDiscussion)
		api.DELETE("/discussions", h.deleteDiscussion)
	}
	return router
}
