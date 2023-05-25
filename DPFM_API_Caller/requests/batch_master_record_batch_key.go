package requests

type BatchMasterRecordBatchKey struct {
	ComponentProduct                 []string `json:"ComponentProduct"`
	StockConfirmationBusinessPartner []int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           []string `json:"StockConfirmationPlant"`
	IsMarkedForDeletion              bool     `json:"IsMarkedForDeletion"`
}
