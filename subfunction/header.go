package subfunction

import (
	api_input_reader "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

func (f *SubFunction) PlannedOrderHederInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PlannedOrderHeder, error) {
	data := make([]*api_processing_data_formatter.PlannedOrderHeder, 0)
	var err error

	processType := psdc.ProcessType

	if processType.PluralitySpec {
		data, err = f.PlannedOrderHederByPluralitySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.PlannedOrderHederByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, xerrors.Errorf("OrderIDの絞り込み（一括登録）に必要な入力パラメータが揃っていません。")
	}

	return data, nil
}

func (f *SubFunction) PlannedOrderHederByPluralitySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PlannedOrderHeder, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToPlannedOrderHederKey()

	mrpArea := sdc.InputParameters.MRPArea
	ownerProductionPlantBusinessPartner := sdc.InputParameters.OwnerProductionPlantBusinessPartner
	ownerProductionPlant := sdc.InputParameters.OwnerProductionPlant
	productInHeader := sdc.InputParameters.ProductInHeader

	dataKey.MRPArea = append(dataKey.MRPArea, *mrpArea...)
	dataKey.OwnerProductionPlantBusinessPartner = append(dataKey.OwnerProductionPlantBusinessPartner, *ownerProductionPlantBusinessPartner...)
	dataKey.OwnerProductionPlant = append(dataKey.OwnerProductionPlant, *ownerProductionPlant...)
	dataKey.ProductInHeader = append(dataKey.ProductInHeader, *productInHeader...)

	repeat1 := strings.Repeat("?,", len(dataKey.MRPArea)-1) + "?"
	for _, v := range dataKey.MRPArea {
		args = append(args, v)
	}
	repeat2 := strings.Repeat("?,", len(dataKey.OwnerProductionPlantBusinessPartner)-1) + "?"
	for _, v := range dataKey.OwnerProductionPlantBusinessPartner {
		args = append(args, v)
	}
	repeat3 := strings.Repeat("?,", len(dataKey.OwnerProductionPlant)-1) + "?"
	for _, v := range dataKey.OwnerProductionPlant {
		args = append(args, v)
	}
	repeat4 := strings.Repeat("?,", len(dataKey.ProductInHeader)-1) + "?"
	for _, v := range dataKey.ProductInHeader {
		args = append(args, v)
	}

	args = append(args, dataKey.PlannedOrderType, dataKey.PlannedOrderIsReleased, dataKey.IsMarkedForDeletion)

	// TODO: BaseUnit除いてる
	rows, err := f.db.Query(
		`SELECT PlannedOrder, PlannedOrderType, Product, ProductDeliverFromParty, ProductDeliverToParty, OriginIssuingPlant,
		OriginIssuingPlantStorageLocation, DestinationReceivingPlant, DestinationReceivingPlantStorageLocation,
		OwnerProductionPlantBusinessPartner, OwnerProductionPlant, OwnerProductionPlantStorageLocation, MRPArea,
		MRPController, PlannedOrderQuantityInBaseUnit, PlannedOrderPlannedScrapQtyInBaseUnit, PlannedOrderOriginIssuingUnit,
		PlannedOrderDestinationReceivingUnit, PlannedOrderOriginIssuingQuantity, PlannedOrderDestinationReceivingQuantity,
		PlannedOrderPlannedStartDate, PlannedOrderPlannedStartTime, PlannedOrderPlannedEndDate, PlannedOrderPlannedEndTime,
		LastChangeDateTime, OrderID, OrderItem, ProductBuyer, ProductSeller, Project, Reservation, ReservationItem,
		PlannedOrderLongText, PlannedOrderIsFixed, PlannedOrderBOMIsFixed, LastScheduledDate, ScheduledBasicEndDate,
		ScheduledBasicEndTime, ScheduledBasicStartDate, ScheduledBasicStartTime, SchedulingType, PlannedOrderIsReleased, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_header_data
		WHERE MRPArea IN ( `+repeat1+` )
		AND OwnerProductionPlantBusinessPartner IN ( `+repeat2+` )
		AND OwnerProductionPlant IN ( `+repeat3+` )
		AND Product IN ( `+repeat4+` )
		AND (PlannedOrderType, PlannedOrderIsReleased, IsMarkedForDeletion) = (?,?,?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToPlannedOrderHeder(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) PlannedOrderHederByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PlannedOrderHeder, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToPlannedOrderHederKey()

	dataKey.MRPAreaTo = sdc.InputParameters.MRPAreaTo
	dataKey.MRPAreaFrom = sdc.InputParameters.MRPAreaFrom
	dataKey.OwnerProductionPlantBusinessPartnerTo = sdc.InputParameters.OwnerProductionPlantBusinessPartnerTo
	dataKey.OwnerProductionPlantBusinessPartnerFrom = sdc.InputParameters.OwnerProductionPlantBusinessPartnerFrom
	dataKey.OwnerProductionPlantTo = sdc.InputParameters.OwnerProductionPlantTo
	dataKey.OwnerProductionPlantFrom = sdc.InputParameters.OwnerProductionPlantFrom
	dataKey.ProductInHeaderTo = sdc.InputParameters.ProductInHeaderTo
	dataKey.ProductInHeaderFrom = sdc.InputParameters.ProductInHeaderFrom

	args = append(args, dataKey.MRPAreaFrom, dataKey.MRPAreaTo,
		dataKey.OwnerProductionPlantBusinessPartnerFrom, dataKey.OwnerProductionPlantBusinessPartnerTo,
		dataKey.OwnerProductionPlantFrom, dataKey.OwnerProductionPlantTo,
		dataKey.ProductInHeaderFrom, dataKey.ProductInHeaderTo,
		dataKey.PlannedOrderType, dataKey.PlannedOrderIsReleased, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT PlannedOrder, PlannedOrderType, Product, ProductDeliverFromParty, ProductDeliverToParty, OriginIssuingPlant,
		OriginIssuingPlantStorageLocation, DestinationReceivingPlant, DestinationReceivingPlantStorageLocation,
		OwnerProductionPlantBusinessPartner, OwnerProductionPlant, OwnerProductionPlantStorageLocation, MRPArea,
		MRPController, PlannedOrderQuantityInBaseUnit, PlannedOrderPlannedScrapQtyInBaseUnit, PlannedOrderOriginIssuingUnit,
		PlannedOrderDestinationReceivingUnit, PlannedOrderOriginIssuingQuantity, PlannedOrderDestinationReceivingQuantity,
		PlannedOrderPlannedStartDate, PlannedOrderPlannedStartTime, PlannedOrderPlannedEndDate, PlannedOrderPlannedEndTime,
		LastChangeDateTime, OrderID, OrderItem, ProductBuyer, ProductSeller, Project, Reservation, ReservationItem,
		PlannedOrderLongText, PlannedOrderIsFixed, PlannedOrderBOMIsFixed, LastScheduledDate, ScheduledBasicEndDate,
		ScheduledBasicEndTime, ScheduledBasicStartDate, ScheduledBasicStartTime, SchedulingType, PlannedOrderIsReleased, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_header_data
		WHERE MRPArea BETWEEN ? AND ?
		AND OwnerProductionPlantBusinessPartner BETWEEN ? AND ?
		AND OwnerProductionPlant BETWEEN ? AND ?
		AND Product BETWEEN ? AND ?
		AND (PlannedOrderType, PlannedOrderIsReleased, IsMarkedForDeletion) = (?,?,?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToPlannedOrderHeder(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) PlannedOrderItemInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PlannedOrderItem, error) {
	data := make([]*api_processing_data_formatter.PlannedOrderItem, 0)
	var err error

	processType := psdc.ProcessType

	if processType.PluralitySpec {
		data, err = f.PlannedOrderItemByPluralitySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.PlannedOrderItemByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, xerrors.Errorf("OrderIDの絞り込み（一括登録）に必要な入力パラメータが揃っていません。")
	}

	return data, nil
}

func (f *SubFunction) PlannedOrderItemByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PlannedOrderItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToPlannedOrderItemKey()

	dataKey.MRPAreaTo = sdc.InputParameters.MRPAreaTo
	dataKey.MRPAreaFrom = sdc.InputParameters.MRPAreaFrom
	dataKey.ProductionPlantBusinessPartnerTo = sdc.InputParameters.ProductionPlantBusinessPartnerTo
	dataKey.ProductionPlantBusinessPartnerFrom = sdc.InputParameters.ProductionPlantBusinessPartnerFrom
	dataKey.ProductionPlantTo = sdc.InputParameters.ProductionPlantTo
	dataKey.ProductionPlantFrom = sdc.InputParameters.ProductionPlantFrom
	dataKey.ProductionPlantStorageLocationTo = sdc.InputParameters.ProductionPlantStorageLocationTo
	dataKey.ProductionPlantStorageLocationFrom = sdc.InputParameters.ProductionPlantStorageLocationFrom
	dataKey.ProductInItemTo = sdc.InputParameters.ProductInItemTo
	dataKey.ProductInItemFrom = sdc.InputParameters.ProductInItemFrom

	args = append(args, dataKey.MRPAreaFrom, dataKey.MRPAreaTo,
		dataKey.ProductionPlantBusinessPartnerFrom, dataKey.ProductionPlantBusinessPartnerTo,
		dataKey.ProductionPlantFrom, dataKey.ProductionPlantTo,
		dataKey.ProductionPlantStorageLocationFrom, dataKey.ProductionPlantStorageLocationTo,
		dataKey.ProductInItemFrom, dataKey.ProductInItemTo,
		dataKey.PlannedOrderIsReleased, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT PlannedOrder, PlannedOrderItem, Product, ProductDeliverFromParty, ProductDeliverToParty, IssuingPlant,
		IssuingPlantStorageLocation, ReceivingPlant, ReceivingPlantStorageLocation, ProductionPlantBusinessPartner,
		ProductionPlant, ProductionPlantStorageLocation, BaseUnit, MRPArea, MRPController, PlannedOrderQuantityInBaseUnit,
		PlannedOrderPlannedScrapQtyInBaseUnit, PlannedOrderIssuingUnit, PlannedOrderReceivingUnit, PlannedOrderIssuingQuantity,
		PlannedOrderReceivingQuantity, PlannedOrderPlannedStartDate, PlannedOrderPlannedStartTime, PlannedOrderPlannedEndDate,
		PlannedOrderPlannedEndTime, LastChangeDateTime, OrderID, OrderItem, ProductBuyer, ProductSeller, Project, Reservation,
		ReservationItem, PlannedOrderLongText, PlannedOrderIsFixed, PlannedOrderBOMIsFixed, LastScheduledDate, ScheduledBasicEndDate,
		ScheduledBasicEndTime, ScheduledBasicStartDate, ScheduledBasicStartTime, SchedulingType, PlannedOrderIsReleased, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_data
		WHERE MRPArea BETWEEN ? AND ?
		AND ProductionPlantBusinessPartner BETWEEN ? AND ?
		AND ProductionPlant BETWEEN ? AND ?
		AND ProductionPlantStorageLocation BETWEEN ? AND ?
		AND Product BETWEEN ? AND ?
		AND (PlannedOrderIsReleased, IsMarkedForDeletion) = (?,?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToPlannedOrderItem(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) PlannedOrderItemByPluralitySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PlannedOrderItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToPlannedOrderItemKey()

	for _, v := range psdc.PlannedOrderHeder {
		dataKey.PlannedOrder = append(dataKey.PlannedOrder, v.PlannedOrder)
	}
	dataKey.MRPArea = append(dataKey.MRPArea, *sdc.InputParameters.MRPArea...)
	dataKey.ProductionPlantBusinessPartner = append(dataKey.ProductionPlantBusinessPartner, *sdc.InputParameters.ProductionPlantBusinessPartner...)
	dataKey.ProductionPlant = append(dataKey.ProductionPlant, *sdc.InputParameters.ProductionPlant...)
	dataKey.ProductionPlantStorageLocation = append(dataKey.ProductionPlantStorageLocation, *sdc.InputParameters.ProductionPlantStorageLocation...)
	dataKey.ProductInItem = append(dataKey.ProductInItem, *sdc.InputParameters.ProductInItem...)

	repeat1 := strings.Repeat("?,", len(dataKey.PlannedOrder)-1) + "?"
	for _, v := range dataKey.PlannedOrder {
		args = append(args, v)
	}
	repeat2 := strings.Repeat("?,", len(dataKey.MRPArea)-1) + "?"
	for _, v := range dataKey.MRPArea {
		args = append(args, v)
	}
	repeat3 := strings.Repeat("?,", len(dataKey.ProductionPlantBusinessPartner)-1) + "?"
	for _, v := range dataKey.ProductionPlantBusinessPartner {
		args = append(args, v)
	}
	repeat4 := strings.Repeat("?,", len(dataKey.ProductionPlant)-1) + "?"
	for _, v := range dataKey.ProductionPlant {
		args = append(args, v)
	}
	repeat5 := strings.Repeat("?,", len(dataKey.ProductionPlantStorageLocation)-1) + "?"
	for _, v := range dataKey.ProductionPlantStorageLocation {
		args = append(args, v)
	}
	repeat6 := strings.Repeat("?,", len(dataKey.ProductInItem)-1) + "?"
	for _, v := range dataKey.ProductInItem {
		args = append(args, v)
	}

	args = append(args, dataKey.PlannedOrderIsReleased, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT PlannedOrder, PlannedOrderItem, Product, ProductDeliverFromParty, ProductDeliverToParty, IssuingPlant,
		IssuingPlantStorageLocation, ReceivingPlant, ReceivingPlantStorageLocation, ProductionPlantBusinessPartner,
		ProductionPlant, ProductionPlantStorageLocation, BaseUnit, MRPArea, MRPController, PlannedOrderQuantityInBaseUnit,
		PlannedOrderPlannedScrapQtyInBaseUnit, PlannedOrderIssuingUnit, PlannedOrderReceivingUnit, PlannedOrderIssuingQuantity,
		PlannedOrderReceivingQuantity, PlannedOrderPlannedStartDate, PlannedOrderPlannedStartTime, PlannedOrderPlannedEndDate,
		PlannedOrderPlannedEndTime, LastChangeDateTime, OrderID, OrderItem, ProductBuyer, ProductSeller, Project, Reservation,
		ReservationItem, PlannedOrderLongText, PlannedOrderIsFixed, PlannedOrderBOMIsFixed, LastScheduledDate, ScheduledBasicEndDate,
		ScheduledBasicEndTime, ScheduledBasicStartDate, ScheduledBasicStartTime, SchedulingType, PlannedOrderIsReleased, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_item_data
		WHERE PlannedOrder IN ( `+repeat1+` )
		AND MRPArea IN ( `+repeat2+` )
		AND ProductionPlantBusinessPartner IN ( `+repeat3+` )
		AND ProductionPlant IN ( `+repeat4+` )
		AND ProductionPlantStorageLocation IN ( `+repeat5+` )
		AND Product IN ( `+repeat6+` )
		AND (PlannedOrderIsReleased, IsMarkedForDeletion) = (?,?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToPlannedOrderItem(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) PlannedOrderComponent(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PlannedOrderComponent, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToPlannedOrderComponentKey()

	for _, v := range psdc.PlannedOrderItem {
		dataKey.PlannedOrder = append(dataKey.PlannedOrder, v.PlannedOrder)
		dataKey.PlannedOrderItem = append(dataKey.PlannedOrderItem, v.PlannedOrderItem)
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.PlannedOrder)-1) + "(?,?)"
	for i := range dataKey.PlannedOrder {
		args = append(args, dataKey.PlannedOrder[i], dataKey.PlannedOrderItem[i])
	}

	args = append(args, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT PlannedOrder, PlannedOrderItem, BOMItem, BOMItemDescription, BillOfMaterialCategory,
		BillOfMaterialItemNumber, BillOfMaterialInternalID, Reservation, ReservationItem, ComponentProduct,
		ComponentProductDeliverFromParty, ComponentProductDeliverToParty, ComponentProductBuyer,
		ComponentProductSeller, ComponentProductRequirementDate, ComponentProductRequiredQuantity, BaseUnit,
		MRPArea, MRPController, StockConfirmationPartnerFunction, StockConfirmationBusinessPartner,
		StockConfirmationPlant, StockConfirmationPlantBatch, StorageLocationForMRP, ComponentWithdrawnQuantity,
		ComponentScrapInPercent, QuantityIsFixed, LastChangeDateTime, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_order_component_data
		WHERE (PlannedOrder, PlannedOrderItem) IN ( `+repeat+` )
		AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToPlannedOrderComponent(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CreationDateHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationDate {
	data := psdc.ConvertToCreationDateHeader(getSystemDate())

	return data
}

func (f *SubFunction) LastChangeDateHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeDate {
	data := psdc.ConvertToLastChangeDateHeader(getSystemDate())

	return data
}

func getSystemDate() string {
	day := time.Now()
	return day.Format("2006-01-02")
}
