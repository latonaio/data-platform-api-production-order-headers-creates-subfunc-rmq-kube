package requests

type PlannedScrapQuantityHeader struct {
	PlannedOrder            int      `json:"PlannedOrder"`
	PlannedOrderItem        int      `json:"PlannedOrderItem"`
	ComponentScrapInPercent float32  `json:"ComponentScrapInPercent"`
	TotalQuantity           float32  `json:"TotalQuantity"`
	PlannedScrapQuantity    *float32 `json:"PlannedScrapQuantity"`
}
