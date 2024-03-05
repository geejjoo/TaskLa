package handler

import (
	"github.com/geejjoo/lamoda_task/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		reservations := api.Group("/reserve")
		{
			reservations.POST("/", h.makeReservation)
			reservations.POST("/cancel", h.cancelReservation)
		}
		inventory := api.Group("/inventory")
		{
			inventory.GET("/:storage", h.storageInformation)
		}
	}

	return router
}
