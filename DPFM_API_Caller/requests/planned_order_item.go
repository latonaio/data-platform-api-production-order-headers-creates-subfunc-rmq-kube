package requests

type PlannedOrderItem struct {
	PlannedOrder                          int      `json:"PlannedOrder"`
	PlannedOrderItem                      int      `json:"PlannedOrderItem"`
	Product                               *string  `json:"Product"`
	ProductDeliverFromParty               *int     `json:"ProductDeliverFromParty"`
	ProductDeliverToParty                 *int     `json:"ProductDeliverToParty"`
	IssuingPlant                          *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation           *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlant                        *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation         *string  `json:"ReceivingPlantStorageLocation"`
	ProductionPlantBusinessPartner        *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                       *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation        *string  `json:"ProductionPlantStorageLocation"`
	BaseUnit                              *string  `json:"BaseUnit"`
	MRPArea                               *string  `json:"MRPArea"`
	MRPController                         *string  `json:"MRPController"`
	PlannedOrderQuantityInBaseUnit        *float32 `json:"PlannedOrderQuantityInBaseUnit"`
	PlannedOrderPlannedScrapQtyInBaseUnit *float32 `json:"PlannedOrderPlannedScrapQtyInBaseUnit"`
	PlannedOrderIssuingUnit               *string  `json:"PlannedOrderIssuingUnit"`
	PlannedOrderReceivingUnit             *string  `json:"PlannedOrderReceivingUnit"`
	PlannedOrderIssuingQuantity           *float32 `json:"PlannedOrderIssuingQuantity"`
	PlannedOrderReceivingQuantity         *float32 `json:"PlannedOrderReceivingQuantity"`
	PlannedOrderPlannedStartDate          *string  `json:"PlannedOrderPlannedStartDate"`
	PlannedOrderPlannedStartTime          *string  `json:"PlannedOrderPlannedStartTime"`
	PlannedOrderPlannedEndDate            *string  `json:"PlannedOrderPlannedEndDate"`
	PlannedOrderPlannedEndTime            *string  `json:"PlannedOrderPlannedEndTime"`
	LastChangeDateTime                    *string  `json:"LastChangeDateTime"`
	OrderID                               *int     `json:"OrderID"`
	OrderItem                             *int     `json:"OrderItem"`
	ProductBuyer                          *int     `json:"ProductBuyer"`
	ProductSeller                         *int     `json:"ProductSeller"`
	Project                               *string  `json:"Project"`
	Reservation                           *int     `json:"Reservation"`
	ReservationItem                       *int     `json:"ReservationItem"`
	PlannedOrderLongText                  *string  `json:"PlannedOrderLongText"`
	PlannedOrderIsFixed                   *bool    `json:"PlannedOrderIsFixed"`
	PlannedOrderBOMIsFixed                *bool    `json:"PlannedOrderBOMIsFixed"`
	LastScheduledDate                     *string  `json:"LastScheduledDate"`
	ScheduledBasicEndDate                 *string  `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime                 *string  `json:"ScheduledBasicEndTime"`
	ScheduledBasicStartDate               *string  `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime               *string  `json:"ScheduledBasicStartTime"`
	SchedulingType                        *string  `json:"SchedulingType"`
	PlannedOrderIsReleased                *bool    `json:"PlannedOrderIsReleased"`
	IsMarkedForDeletion                   *bool    `json:"IsMarkedForDeletion"`
}
