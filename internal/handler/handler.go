package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pelicanch1k/rest-api/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	products := router.Group("/products/")
	{
		products.POST("", h.create)
		products.GET("", h.get)

		products.PUT(":id", h.update)
		products.DELETE(":id", h.delete)
		products.GET(":id", h.getAll)
	}

	return router
}
