package requests

type ItemIsConfirmed struct {
	PlannedOrder        int  `json:"PlannedOrder"`
	PlannedOrderItem    int  `json:"PlannedOrderItem"`
	StockIsFullyChecked bool `json:"StockIsFullyChecked"`
	ItemIsConfirmed     bool `json:"ItemIsConfirmed"`
}
