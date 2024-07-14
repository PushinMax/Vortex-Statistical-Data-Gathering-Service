package repository

import (
	"github.com/jmoiron/sqlx"
	"vortex-statistical-data-gathering-service/internal/utils"
)

type MyApi interface {
	GetOrderBook(exchange_name, pair string) ([]*utils.DepthOrder, error)
	SaveOrderBook(exchange_name, pair string, orderBook []*utils.DepthOrder) error
	GetOrderHistory(client *utils.Client) ([]*utils.HistoryOrder, error)
	SaveOrder(client *utils.Client, order *utils.HistoryOrder) error
}

type Repository struct {
	MyApi
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{MyApi: newPostgresImpl(db)}
}
