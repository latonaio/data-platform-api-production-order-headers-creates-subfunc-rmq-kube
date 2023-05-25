package requests

type ItemIsPartiallyConfirmed struct {
	PlannedOrder                    int     `json:"PlannedOrder"`
	PlannedOrderItem                int     `json:"PlannedOrderItem"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	ItemIsPartiallyConfirmed        bool    `json:"ItemIsPartiallyConfirmed"`
}
