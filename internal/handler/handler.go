package handler

import (
	"github.com/gin-gonic/gin"
	"vortex-statistical-data-gathering-service/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/get_order_book", h.GetOrderBook)
		api.GET("/get_order_history", h.GetOrderHistory)

		api.POST("/save_order_book", h.SaveOrderBook)
		api.POST("/save_order", h.SaveOrder)
	}
	return router
}
