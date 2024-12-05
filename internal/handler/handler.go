package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pelicanch1k/ProductGatewayAPI/internal/service"
	"github.com/pelicanch1k/ProductGatewayAPI/pkg/logging"
)

type Handler struct {
	services *service.Service
	logger   *logging.Logger
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services, logger: logging.GetLogger()}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth/", h.logging)
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	products := router.Group("/products/", h.logging, h.userIdentity)
	{
		products.POST("", h.create)
		products.GET("", h.getAll)

		products.PUT(":id", h.update)
		products.DELETE(":id", h.delete)
		products.GET(":id", h.get)
	}

	return router
}
