package api_processing_data_formatter

type SDC struct {
	MetaData                        *MetaData                        `json:"MetaData"`
	ProcessType                     *ProcessType                     `json:"ProcessType"`
	PlannedOrderHeader              []*PlannedOrderHeader            `json:"PlannedOrderHeader"`
	PlannedOrderItem                []*PlannedOrderItem              `json:"PlannedOrderItem"`
	PlannedOrderComponent           []*PlannedOrderComponent         `json:"PlannedOrderComponent"`
	PlannedScrapQuantityHeader      []*PlannedScrapQuantityHeader    `json:"PlannedScrapQuantityHeader"`
	ExecuteProductAvailabilityCheck *ExecuteProductAvailabilityCheck `json:"ExecuteProductAvailabilityCheck"`
	ProductMasterBPPlant            []*ProductMasterBPPlant          `json:"ProductMasterBPPlant"`
	BatchMasterRecordBatch          []*BatchMasterRecordBatch        `json:"BatchMasterRecordBatch"`
	StockConfirmation               []*StockConfirmation             `json:"StockConfirmation"`
	ItemIsPartiallyConfirmed        []*ItemIsPartiallyConfirmed      `json:"ItemIsPartiallyConfirmed"`
	ItemIsConfirmed                 []*ItemIsConfirmed               `json:"ItemIsConfirmed"`
	TotalQuantity                   []*TotalQuantity                 `json:"TotalQuantity"`
	HeaderIsPartiallyConfirmed      *HeaderIsPartiallyConfirmed      `json:"HeaderIsPartiallyConfirmed"`
	HeaderIsConfirmed               *HeaderIsConfirmed               `json:"HeaderIsConfirmed"`
	TotalQuantityHeader             *TotalQuantityHeader             `json:"TotalQuantityHeader"`
	CreationDateHeader              *CreationDate                    `json:"CreationDateHeader"`
	LastChangeDateHeader            *LastChangeDate                  `json:"LastChangeDateHeader"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type ProcessType struct {
	BulkProcess       bool `json:"BulkProcess"`
	IndividualProcess bool `json:"IndividualProcess"`
	PluralitySpec     bool `json:"PluralitySpec"`
	RangeSpec         bool `json:"RangeSpec"`
}

// Header
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

type PlannedOrderComponentKey struct {
	PlannedOrder        []int `json:"PlannedOrder"`
	PlannedOrderItem    []int `json:"PlannedOrderItem"`
	IsMarkedForDeletion bool  `json:"IsMarkedForDeletion"`
}

type PlannedOrderComponent struct {
	PlannedOrder                     int      `json:"PlannedOrder"`
	PlannedOrderItem                 int      `json:"PlannedOrderItem"`
	BillOfMaterial                   *int     `json:"BillOfMaterial"`
	BOMItem                          *int     `json:"BOMItem"`
	Operations                       *int     `json:"Operations"`
	OperationsItem                   *int     `json:"OperationsItem"`
	Reservation                      *int     `json:"Reservation"`
	ReservationItem                  *int     `json:"ReservationItem"`
	ComponentProduct                 *string  `json:"ComponentProduct"`
	ComponentProductDeliverFromParty *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverToParty   *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductBuyer            *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller           *int     `json:"ComponentProductSeller"`
	ComponentProductRequirementDate  *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime  *string  `json:"ComponentProductRequirementTime"`
	ComponentProductRequiredQuantity *float32 `json:"ComponentProductRequiredQuantity"`
	ComponentProductBusinessPartner  *int     `json:"ComponentProductBusinessPartner"`
	BaseUnit                         *string  `json:"BaseUnit"`
	MRPArea                          *string  `json:"MRPArea"`
	MRPController                    *string  `json:"MRPController"`
	StockConfirmationPartnerFunction *string  `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch      *string  `json:"StockConfirmationPlantBatch"`
	StorageLocationForMRP            *string  `json:"StorageLocationForMRP"`
	ComponentWithdrawnQuantity       *float32 `json:"ComponentWithdrawnQuantity"`
	ComponentScrapInPercent          *float32 `json:"ComponentScrapInPercent"`
	OperationScrapInPercent          *float32 `json:"OperationScrapInPercent"`
	QuantityIsFixed                  *bool    `json:"QuantityIsFixed"`
	LastChangeDateTime               *string  `json:"LastChangeDateTime"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
}

type PlannedScrapQuantityHeader struct {
	PlannedOrder            int      `json:"PlannedOrder"`
	PlannedOrderItem        int      `json:"PlannedOrderItem"`
	ComponentScrapInPercent float32  `json:"ComponentScrapInPercent"`
	TotalQuantity           float32  `json:"TotalQuantity"`
	PlannedScrapQuantity    *float32 `json:"PlannedScrapQuantity"`
}

// Item
type ExecuteProductAvailabilityCheck struct {
	ExecuteProductAvailabilityCheck bool `json:"ExecuteProductAvailabilityCheck"`
}

type ProductMasterBPPlantKey struct {
	ComponentProduct                 []string `json:"ComponentProduct"`
	StockConfirmationBusinessPartner []int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           []string `json:"StockConfirmationPlant"`
}

type ProductMasterBPPlant struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	IsBatchManagementRequired *bool   `json:"IsBatchManagementRequired"`
	BatchManagementPolicy     *string `json:"BatchManagementPolicy"`
}

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

type StockConfirmation struct {
	PlannedOrder                    int     `json:"PlannedOrder"`
	PlannedOrderItem                int     `json:"PlannedOrderItem"`
	BusinessPartner                 int     `json:"BusinessPartner"`
	Product                         string  `json:"Product"`
	Plant                           string  `json:"Plant"`
	Batch                           string  `json:"Batch"`
	RequestedQuantity               float32 `json:"RequestedQuantity"`
	ProductStockAvailabilityDate    string  `json:"ProductStockAvailabilityDate"`
	OrderID                         int     `json:"OrderID"`
	OrderItem                       int     `json:"OrderItem"`
	Project                         string  `json:"Project"`
	InventoryStockType              string  `json:"InventoryStockType"`
	InventorySpecialStockType       string  `json:"InventorySpecialStockType"`
	AvailableProductStock           float32 `json:"AvailableProductStock"`
	CheckedQuantity                 float32 `json:"CheckedQuantity"`
	CheckedDate                     string  `json:"CheckedDate"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
	Suffix                          string  `json:"Suffix"`
	StockConfirmationIsLotUnit      bool    `json:"StockConfirmationIsLotUnit"`
	StockConfirmationIsOrdinary     bool    `json:"StockConfirmationIsOrdinary"`
}

type ProductAvailabilityCheck struct {
	ConnectionKey                 string `json:"connection_key"`
	Result                        bool   `json:"result"`
	RedisKey                      string `json:"redis_key"`
	Filepath                      string `json:"filepath"`
	APIStatusCode                 int    `json:"api_status_code"`
	RuntimeSessionID              string `json:"runtime_session_id"`
	BusinessPartnerID             *int   `json:"business_partner"`
	ServiceLabel                  string `json:"service_label"`
	ProductStockAvailabilityCheck struct {
		Product                         *string  `json:"Product"`
		BusinessPartner                 *int     `json:"BusinessPartner"`
		Plant                           *string  `json:"Plant"`
		StorageLocation                 *string  `json:"StorageLocation"`
		StorageBin                      *string  `json:"StorageBin"`
		Batch                           *string  `json:"Batch"`
		RequestedQuantity               *float32 `json:"RequestedQuantity"`
		ProductStockAvailabilityDate    *string  `json:"ProductStockAvailabilityDate"`
		InventoryStockType              *string  `json:"InventoryStockType"`
		InventorySpecialStockType       *string  `json:"InventorySpecialStockType"`
		AvailableProductStock           *float32 `json:"AvailableProductStock"`
		CheckedQuantity                 *float32 `json:"CheckedQuantity"`
		CheckedDate                     *string  `json:"CheckedDate"`
		OpenConfirmedQuantityInBaseUnit *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
		StockIsFullyChecked             *bool    `json:"StockIsFullyChecked"`
		Suffix                          *string  `json:"Suffix"`
	} `json:"ProductStockAvailabilityCheck"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	ProductStockCode string   `json:"product_stock_code"`
	Deleted          bool     `json:"deleted"`
}

type BatchMasterRecordBatchKey struct {
	ComponentProduct                 []string `json:"ComponentProduct"`
	StockConfirmationBusinessPartner []int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           []string `json:"StockConfirmationPlant"`
	IsMarkedForDeletion              bool     `json:"IsMarkedForDeletion"`
}

type BatchMasterRecordBatch struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	Batch               string `json:"Batch"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityStartTime   string `json:"ValidityStartTime"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	ValidityEndTime     string `json:"ValidityEndTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ItemIsPartiallyConfirmed struct {
	PlannedOrder                    int     `json:"PlannedOrder"`
	PlannedOrderItem                int     `json:"PlannedOrderItem"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	ItemIsPartiallyConfirmed        bool    `json:"ItemIsPartiallyConfirmed"`
}

type ItemIsConfirmed struct {
	PlannedOrder        int  `json:"PlannedOrder"`
	PlannedOrderItem    int  `json:"PlannedOrderItem"`
	StockIsFullyChecked bool `json:"StockIsFullyChecked"`
	ItemIsConfirmed     bool `json:"ItemIsConfirmed"`
}

type TotalQuantity struct {
	PlannedOrder     int      `json:"PlannedOrder"`
	PlannedOrderItem int      `json:"PlannedOrderItem"`
	TotalQuantity    *float32 `json:"TotalQuantity"`
}

// Headerの集計項目等の計算とセット
type HeaderIsPartiallyConfirmed struct {
	HeaderIsPartiallyConfirmed *bool `json:"HeaderIsPartiallyConfirmed"`
}

type HeaderIsConfirmed struct {
	HeaderIsConfirmed *bool `json:"HeaderIsConfirmed"`
}

type TotalQuantityHeader struct {
	TotalQuantity *float32 `json:"TotalQuantity"`
}

// 日付等の処理
type CreationDate struct {
	CreationDate string `json:"CreationDate"`
}

type LastChangeDate struct {
	LastChangeDate string `json:"LastChangeDate"`
}
