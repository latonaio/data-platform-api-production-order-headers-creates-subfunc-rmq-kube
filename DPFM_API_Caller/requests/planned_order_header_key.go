package requests

type PlannedOrderHeaderKey struct {
	MRPArea                                 []*string `json:"MRPArea"`
	MRPAreaTo                               *string   `json:"MRPAreaTo"`
	MRPAreaFrom                             *string   `json:"MRPAreaFrom"`
	OwnerProductionPlantBusinessPartner     []*int    `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerTo   *int      `json:"OwnerProductionPlantBusinessPartnerTo"`
	OwnerProductionPlantBusinessPartnerFrom *int      `json:"OwnerProductionPlantBusinessPartnerFrom"`
	OwnerProductionPlant                    []*string `json:"OwnerProductionPlant"`
	OwnerProductionPlantTo                  *string   `json:"OwnerProductionPlantTo"`
	OwnerProductionPlantFrom                *string   `json:"OwnerProductionPlantFrom"`
	ProductInHeader                         []*string `json:"ProductInHeader"`
	ProductInHeaderTo                       *string   `json:"ProductInHeaderTo"`
	ProductInHeaderFrom                     *string   `json:"ProductInHeaderFrom"`
	PlannedOrderType                        string    `json:"PlannedOrderType"`
	PlannedOrderIsReleased                  bool      `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                     bool      `json:"IsMarkedForDeletion"`
}
