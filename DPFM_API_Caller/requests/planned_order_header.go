package requests

type PlannedOrderHeader struct {
	PlannedOrder                             int      `json:"PlannedOrder"`
	PlannedOrderType                         *string  `json:"PlannedOrderType"`
	Product                                  *string  `json:"Product"`
	ProductDeliverFromParty                  *int     `json:"ProductDeliverFromParty"`
	ProductDeliverToParty                    *int     `json:"ProductDeliverToParty"`
	OriginIssuingPlant                       *string  `json:"OriginIssuingPlant"`
	OriginIssuingPlantStorageLocation        *string  `json:"OriginIssuingPlantStorageLocation"`
	DestinationReceivingPlant                *string  `json:"DestinationReceivingPlant"`
	DestinationReceivingPlantStorageLocation *string  `json:"DestinationReceivingPlantStorageLocation"`
	OwnerProductionPlantBusinessPartner      *int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     *string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantStorageLocation      *string  `json:"OwnerProductionPlantStorageLocation"`
	MRPArea                                  *string  `json:"MRPArea"`
	MRPController                            *string  `json:"MRPController"`
	PlannedOrderQuantityInBaseUnit           *float32 `json:"PlannedOrderQuantityInBaseUnit"`
	PlannedOrderPlannedScrapQtyInBaseUnit    *float32 `json:"PlannedOrderPlannedScrapQtyInBaseUnit"`
	PlannedOrderOriginIssuingUnit            *string  `json:"PlannedOrderOriginIssuingUnit"`
	PlannedOrderDestinationReceivingUnit     *string  `json:"PlannedOrderDestinationReceivingUnit"`
	PlannedOrderOriginIssuingQuantity        *float32 `json:"PlannedOrderOriginIssuingQuantity"`
	PlannedOrderDestinationReceivingQuantity *float32 `json:"PlannedOrderDestinationReceivingQuantity"`
	PlannedOrderPlannedStartDate             *string  `json:"PlannedOrderPlannedStartDate"`
	PlannedOrderPlannedStartTime             *string  `json:"PlannedOrderPlannedStartTime"`
	PlannedOrderPlannedEndDate               *string  `json:"PlannedOrderPlannedEndDate"`
	PlannedOrderPlannedEndTime               *string  `json:"PlannedOrderPlannedEndTime"`
	LastChangeDateTime                       *string  `json:"LastChangeDateTime"`
	OrderID                                  *int     `json:"OrderID"`
	OrderItem                                *int     `json:"OrderItem"`
	ProductBuyer                             *int     `json:"ProductBuyer"`
	ProductSeller                            *int     `json:"ProductSeller"`
	Project                                  *string  `json:"Project"`
	Reservation                              *int     `json:"Reservation"`
	ReservationItem                          *int     `json:"ReservationItem"`
	PlannedOrderLongText                     *string  `json:"PlannedOrderLongText"`
	PlannedOrderIsFixed                      *bool    `json:"PlannedOrderIsFixed"`
	PlannedOrderBOMIsFixed                   *bool    `json:"PlannedOrderBOMIsFixed"`
	LastScheduledDate                        *string  `json:"LastScheduledDate"`
	ScheduledBasicEndDate                    *string  `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime                    *string  `json:"ScheduledBasicEndTime"`
	ScheduledBasicStartDate                  *string  `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime                  *string  `json:"ScheduledBasicStartTime"`
	SchedulingType                           *string  `json:"SchedulingType"`
	PlannedOrderIsReleased                   *bool    `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}
