package requests

type PlannedOrderItemKey struct {
	PlannedOrder                       []int     `json:"PlannedOrder"`
	MRPArea                            []*string `json:"MRPArea"`
	MRPAreaTo                          *string   `json:"MRPAreaTo"`
	MRPAreaFrom                        *string   `json:"MRPAreaFrom"`
	ProductionPlantBusinessPartner     []*int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerTo   *int      `json:"ProductionPlantBusinessPartnerTo"`
	ProductionPlantBusinessPartnerFrom *int      `json:"ProductionPlantBusinessPartnerFrom"`
	ProductionPlant                    []*string `json:"ProductionPlant"`
	ProductionPlantTo                  *string   `json:"ProductionPlantTo"`
	ProductionPlantFrom                *string   `json:"ProductionPlantFrom"`
	ProductionPlantStorageLocation     []*string `json:"ProductionPlantStorageLocation"`
	ProductionPlantStorageLocationTo   *string   `json:"ProductionPlantStorageLocationTo"`
	ProductionPlantStorageLocationFrom *string   `json:"ProductionPlantStorageLocationFrom"`
	ProductInItem                      []*string `json:"ProductInItem"`
	ProductInItemTo                    *string   `json:"ProductInItemTo"`
	ProductInItemFrom                  *string   `json:"ProductInItemFrom"`
	PlannedOrderIsReleased             bool      `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                bool      `json:"IsMarkedForDeletion"`
}
