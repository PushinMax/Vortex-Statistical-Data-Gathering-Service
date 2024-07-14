package utils

import "time"

type GetOrderBookRequest struct {
	ExchangeName string `json:"exchange_name"`
	Pair         string `json:"pair"`
}

type SaveOrderBookRequest struct {
	ExchangeName string       `json:"exchange_name"`
	Pair         string       `json:"pair"`
	OrderBook    []DepthOrder `json:"orderBook"`
}

type GetOrderHistoryRequest struct {
	Client Client `json:"client"`
}

type SaveOrderRequest struct {
	Client Client       `json:"client"`
	Order  HistoryOrder `json:"order"`
}

type DepthOrder struct {
	Price   float64 `json:"Price"`
	BaseQty float64 `json:"BaseQty"`
}

type HistoryOrder struct {
	ClientName          string    `json:"client_name"`
	ExchangeName        string    `json:"exchange_name"`
	Label               string    `json:"label"`
	Pair                string    `json:"pair"`
	Side                string    `json:"side"`
	Type                string    `json:"type"`
	BaseQty             float64   `json:"base_Qty"`
	Price               float64   `json:"price"`
	AlgorithmNamePlaced string    `json:"algorithm_name_placed"`
	LowestSellPrc       float64   `json:"lowest_sell_prc"`
	HighestBuyPrc       float64   `json:"highest_buy_prc"`
	CommissionQuoteQty  float64   `json:"commission_quote_qty"`
	TimePlaced          time.Time `json:"time_placed"`
}

type Client struct {
	ClientName   string `json:"client_name"`
	ExchangeName string `json:"exchange_name"`
	Label        string `json:"label"`
	Pair         string `json:"pair"`
}
