package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vortex-statistical-data-gathering-service/internal/utils"
)

type postgresImpl struct {
	db *sqlx.DB
}

func newPostgresImpl(db *sqlx.DB) *postgresImpl {
	return &postgresImpl{db: db}
}

func (r *postgresImpl) GetOrderBook(exchange_name, pair string) ([]*utils.DepthOrder, error) {
	var id int64
	query := fmt.Sprintf("SELECT id FROM orderBook WHERE exchange = $1 AND pair = $2")
	row := r.db.QueryRow(query, exchange_name, pair)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	var asks []*utils.DepthOrder
	query = fmt.Sprintf("SELECT price, baseQty FROM asks WHERE id_book = $1")
	row = r.db.QueryRow(query, id)
	if err := row.Scan(asks); err != nil {
		return nil, err
	}
	return asks, nil
}

func (r *postgresImpl) SaveOrderBook(exchange_name, pair string, orderBook []*utils.DepthOrder) error {
	var id int64
	query := fmt.Sprint("SELECT id FROM orderBook WHERE exchange = $1 AND pair = $2")
	row := r.db.QueryRow(query, exchange_name, pair)
	if err := row.Scan(&id); err != nil {
		return err
	}

	query = fmt.Sprint("INSERT INTO bids (id_book, price, baseQty) values ($1, $2, $3)")
	for _, v := range orderBook {
		row = r.db.QueryRow(query, id, v.Price, v.BaseQty)
		if err := row.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (r *postgresImpl) GetOrderHistory(client *utils.Client) ([]*utils.HistoryOrder, error) {
	var history []*utils.HistoryOrder
	query := fmt.Sprint("SELECT * FROM orderHistory WHERE client_name = $1 AND exchange_name = $2 AND label = $3 AND pair = $4")
	row := r.db.QueryRow(query, client.ClientName, client.ExchangeName, client.Label, client.Pair)
	if err := row.Scan(history); err != nil {
		return nil, err
	}
	return history, nil
}

func (r *postgresImpl) SaveOrder(client *utils.Client, order *utils.HistoryOrder) error {
	query := fmt.Sprint("INSERT INTO orderHistory (client_name, exchange_name, label, pair, side, type, base_qty, price, algorithm_name_placed, lowest_sell_prc, highest_buy_prc, commission_quote_qty, time_placed) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)")
	row := r.db.QueryRow(
		query,
		order.ClientName, order.ExchangeName, order.Label, order.Pair,
		order.Side, order.Type, order.BaseQty, order.Price, order.AlgorithmNamePlaced,
		order.LowestSellPrc, order.HighestBuyPrc, order.CommissionQuoteQty, order.TimePlaced)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}
