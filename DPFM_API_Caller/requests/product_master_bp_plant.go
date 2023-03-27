package requests

type ProductMasterBPPlant struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	IsBatchManagementRequired *bool   `json:"IsBatchManagementRequired"`
	BatchManagementPolicy     *string `json:"BatchManagementPolicy"`
}
