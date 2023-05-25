package requests

type StockConfirmationKey struct {
	PlannedOrder                                 int     `json:"PlannedOrder"`
	PlannedOrderItem                             int     `json:"PlannedOrderItem"`
	Product                                      string  `json:"Product"`
	StockConfirmationBusinessPartner             int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	PlannedOrderIssuingQuantity                  float32 `json:"PlannedOrderIssuingQuantity"`
	ComponentProductRequirementDate              string  `json:"ComponentProductRequirementDate"`
	StockConfirmationIsLotUnit                   bool    `json:"StockConfirmationIsLotUnit"`
	StockConfirmationIsOrdinary                  bool    `json:"StockConfirmationIsOrdinary"`
}
