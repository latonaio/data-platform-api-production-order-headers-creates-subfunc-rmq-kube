package requests

type ProductMasterBPPlantKey struct {
	ComponentProduct               []string `json:"ComponentProduct"`
	ProductionPlantBusinessPartner []int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                []string `json:"ProductionPlant"`
}
