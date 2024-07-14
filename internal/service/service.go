package service

import (
	"vortex-statistical-data-gathering-service/internal/repository"
	"vortex-statistical-data-gathering-service/internal/utils"
)

type MyApi interface {
	GetOrderBook(exchange_name, pair string) ([]*utils.DepthOrder, error)
	SaveOrderBook(exchange_name, pair string, orderBook []*utils.DepthOrder) error
	GetOrderHistory(client *utils.Client) ([]*utils.HistoryOrder, error)
	SaveOrder(client *utils.Client, order *utils.HistoryOrder) error
}

type Service struct {
	MyApi
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		MyApi: NewApiService(repos),
	}
}
