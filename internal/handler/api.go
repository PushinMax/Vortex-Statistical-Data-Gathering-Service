package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vortex-statistical-data-gathering-service/internal/utils"
)

func (h *Handler) GetOrderBook(c *gin.Context) {
	var request utils.GetOrderBookRequest
	if err := c.BindJSON(request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	asks, err := h.services.MyApi.GetOrderBook(request.ExchangeName, request.Pair)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, asks)
}

func (h *Handler) SaveOrderBook(c *gin.Context) {
	var request utils.SaveOrderBookRequest
	if err := c.BindJSON(request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.MyApi.SaveOrderBook(request.ExchangeName, request.Pair, request.OrderBook)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetOrderHistory(c *gin.Context) {
	var request utils.GetOrderHistoryRequest
	if err := c.BindJSON(request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	orders, err := h.services.MyApi.GetOrderHistory(&request.Client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) SaveOrder(c *gin.Context) {
	var request utils.SaveOrderRequest
	if err := c.BindJSON(request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.MyApi.SaveOrder(&request.Client, &request.Order)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
