package service

import (
	"vortex-statistical-data-gathering-service/internal/repository"
	"vortex-statistical-data-gathering-service/internal/utils"
)

type ApiService struct {
	repo repository.MyApi
}

func NewApiService(repo repository.MyApi) *ApiService {
	return &ApiService{repo: repo}
}

func (s *ApiService) GetOrderBook(exchange_name, pair string) ([]*utils.DepthOrder, error) {
	return s.repo.GetOrderBook(exchange_name, pair)
}
func (s *ApiService) SaveOrderBook(exchange_name, pair string, orderBook []*utils.DepthOrder) error {
	return s.repo.SaveOrderBook(exchange_name, pair, orderBook)
}

func (s *ApiService) GetOrderHistory(client *utils.Client) ([]*utils.HistoryOrder, error) {
	return s.repo.GetOrderHistory(client)
}

func (s *ApiService) SaveOrder(client *utils.Client, order *utils.HistoryOrder) error {
	return s.repo.SaveOrder(client, order)
}
