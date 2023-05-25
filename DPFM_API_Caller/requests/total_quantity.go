package requests

type TotalQuantity struct {
	PlannedOrder     int      `json:"PlannedOrder"`
	PlannedOrderItem int      `json:"PlannedOrderItem"`
	TotalQuantity    *float32 `json:"TotalQuantity"`
}
