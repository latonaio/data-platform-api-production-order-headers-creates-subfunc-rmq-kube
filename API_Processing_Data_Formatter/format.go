package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-production-order-headers-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"

	"golang.org/x/xerrors"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}

	data := pm
	res := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &res
}

func (psdc *SDC) ConvertToProcessType() *ProcessType {
	pm := &requests.ProcessType{}
	data := pm

	processType := ProcessType{
		BulkProcess:       data.BulkProcess,
		IndividualProcess: data.IndividualProcess,
		PluralitySpec:     data.PluralitySpec,
		RangeSpec:         data.RangeSpec,
	}

	return &processType
}

// Header
func (psdc *SDC) ConvertToPlannedOrderHederKey() *PlannedOrderHederKey {
	pm := &requests.PlannedOrderHederKey{
		PlannedOrderType:       "PRD",
		PlannedOrderIsReleased: true,
		IsMarkedForDeletion:    false,
	}

	data := pm
	res := PlannedOrderHederKey{
		MRPArea:                             data.MRPArea,
		MRPAreaTo:                           data.MRPAreaTo,
		MRPAreaFrom:                         data.MRPAreaFrom,
		OwnerProductionPlantBusinessPartner: data.OwnerProductionPlantBusinessPartner,
		OwnerProductionPlant:                data.OwnerProductionPlant,
		OwnerProductionPlantTo:              data.OwnerProductionPlantTo,
		OwnerProductionPlantFrom:            data.OwnerProductionPlantFrom,
		ProductInHeader:                     data.ProductInHeader,
		PlannedOrderType:                    data.PlannedOrderType,
		PlannedOrderIsReleased:              data.PlannedOrderIsReleased,
		IsMarkedForDeletion:                 data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToPlannedOrderHeder(rows *sql.Rows) ([]*PlannedOrderHeder, error) {
	defer rows.Close()
	res := make([]*PlannedOrderHeder, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PlannedOrderHeder{}

		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderType,
			&pm.Product,
			&pm.ProductDeliverFromParty,
			&pm.ProductDeliverToParty,
			&pm.OriginIssuingPlant,
			&pm.OriginIssuingPlantStorageLocation,
			&pm.DestinationReceivingPlant,
			&pm.DestinationReceivingPlantStorageLocation,
			&pm.OwnerProductionPlantBusinessPartner,
			&pm.OwnerProductionPlant,
			&pm.OwnerProductionPlantStorageLocation,
			// &pm.BaseUnit,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.PlannedOrderQuantityInBaseUnit,
			&pm.PlannedOrderPlannedScrapQtyInBaseUnit,
			&pm.PlannedOrderOriginIssuingUnit,
			&pm.PlannedOrderDestinationReceivingUnit,
			&pm.PlannedOrderOriginIssuingQuantity,
			&pm.PlannedOrderDestinationReceivingQuantity,
			&pm.PlannedOrderPlannedStartDate,
			&pm.PlannedOrderPlannedStartTime,
			&pm.PlannedOrderPlannedEndDate,
			&pm.PlannedOrderPlannedEndTime,
			&pm.LastChangeDateTime,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ProductBuyer,
			&pm.ProductSeller,
			&pm.Project,
			&pm.Reservation,
			&pm.ReservationItem,
			&pm.PlannedOrderLongText,
			&pm.PlannedOrderIsFixed,
			&pm.PlannedOrderBOMIsFixed,
			&pm.LastScheduledDate,
			&pm.ScheduledBasicEndDate,
			&pm.ScheduledBasicEndTime,
			&pm.ScheduledBasicStartDate,
			&pm.ScheduledBasicStartTime,
			&pm.SchedulingType,
			&pm.PlannedOrderIsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}
		data := pm
		res = append(res, &PlannedOrderHeder{
			PlannedOrder:                             data.PlannedOrder,
			PlannedOrderType:                         data.PlannedOrderType,
			Product:                                  data.Product,
			ProductDeliverFromParty:                  data.ProductDeliverFromParty,
			ProductDeliverToParty:                    data.ProductDeliverToParty,
			OriginIssuingPlant:                       data.OriginIssuingPlant,
			OriginIssuingPlantStorageLocation:        data.OriginIssuingPlantStorageLocation,
			DestinationReceivingPlant:                data.DestinationReceivingPlant,
			DestinationReceivingPlantStorageLocation: data.DestinationReceivingPlantStorageLocation,
			OwnerProductionPlantBusinessPartner:      data.OwnerProductionPlantBusinessPartner,
			OwnerProductionPlant:                     data.OwnerProductionPlant,
			OwnerProductionPlantStorageLocation:      data.OwnerProductionPlantStorageLocation,
			// BaseUnit:                                 data.BaseUnit,
			MRPArea:                                  data.MRPArea,
			MRPController:                            data.MRPController,
			PlannedOrderQuantityInBaseUnit:           data.PlannedOrderQuantityInBaseUnit,
			PlannedOrderPlannedScrapQtyInBaseUnit:    data.PlannedOrderPlannedScrapQtyInBaseUnit,
			PlannedOrderOriginIssuingUnit:            data.PlannedOrderOriginIssuingUnit,
			PlannedOrderDestinationReceivingUnit:     data.PlannedOrderDestinationReceivingUnit,
			PlannedOrderOriginIssuingQuantity:        data.PlannedOrderOriginIssuingQuantity,
			PlannedOrderDestinationReceivingQuantity: data.PlannedOrderDestinationReceivingQuantity,
			PlannedOrderPlannedStartDate:             data.PlannedOrderPlannedStartDate,
			PlannedOrderPlannedStartTime:             data.PlannedOrderPlannedStartTime,
			PlannedOrderPlannedEndDate:               data.PlannedOrderPlannedEndDate,
			PlannedOrderPlannedEndTime:               data.PlannedOrderPlannedEndTime,
			LastChangeDateTime:                       data.LastChangeDateTime,
			OrderID:                                  data.OrderID,
			OrderItem:                                data.OrderItem,
			ProductBuyer:                             data.ProductBuyer,
			ProductSeller:                            data.ProductSeller,
			Project:                                  data.Project,
			Reservation:                              data.Reservation,
			ReservationItem:                          data.ReservationItem,
			PlannedOrderLongText:                     data.PlannedOrderLongText,
			PlannedOrderIsFixed:                      data.PlannedOrderIsFixed,
			PlannedOrderBOMIsFixed:                   data.PlannedOrderBOMIsFixed,
			LastScheduledDate:                        data.LastScheduledDate,
			ScheduledBasicEndDate:                    data.ScheduledBasicEndDate,
			ScheduledBasicEndTime:                    data.ScheduledBasicEndTime,
			ScheduledBasicStartDate:                  data.ScheduledBasicStartDate,
			ScheduledBasicStartTime:                  data.ScheduledBasicStartTime,
			SchedulingType:                           data.SchedulingType,
			PlannedOrderIsReleased:                   data.PlannedOrderIsReleased,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_planned_order_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToPlannedOrderItemKey() *PlannedOrderItemKey {
	pm := &requests.PlannedOrderItemKey{
		PlannedOrderIsReleased: true,
		IsMarkedForDeletion:    false,
	}

	data := pm
	res := PlannedOrderItemKey{
		PlannedOrder:                       data.PlannedOrder,
		MRPArea:                            data.MRPArea,
		MRPAreaTo:                          data.MRPAreaTo,
		MRPAreaFrom:                        data.MRPAreaFrom,
		ProductionPlantBusinessPartner:     data.ProductionPlantBusinessPartner,
		ProductionPlant:                    data.ProductionPlant,
		ProductionPlantTo:                  data.ProductionPlantTo,
		ProductionPlantFrom:                data.ProductionPlantFrom,
		ProductionPlantStorageLocation:     data.ProductionPlantStorageLocation,
		ProductionPlantStorageLocationTo:   data.ProductionPlantStorageLocationTo,
		ProductionPlantStorageLocationFrom: data.ProductionPlantStorageLocationFrom,
		ProductInItem:                      data.ProductInItem,
		ProductInItemTo:                    data.ProductInItemTo,
		ProductInItemFrom:                  data.ProductInItemFrom,
		PlannedOrderIsReleased:             data.PlannedOrderIsReleased,
		IsMarkedForDeletion:                data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToPlannedOrderItem(rows *sql.Rows) ([]*PlannedOrderItem, error) {
	defer rows.Close()
	res := make([]*PlannedOrderItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PlannedOrderItem{}

		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.Product,
			&pm.ProductDeliverFromParty,
			&pm.ProductDeliverToParty,
			&pm.IssuingPlant,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantStorageLocation,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.BaseUnit,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.PlannedOrderQuantityInBaseUnit,
			&pm.PlannedOrderPlannedScrapQtyInBaseUnit,
			&pm.PlannedOrderIssuingUnit,
			&pm.PlannedOrderReceivingUnit,
			&pm.PlannedOrderIssuingQuantity,
			&pm.PlannedOrderReceivingQuantity,
			&pm.PlannedOrderPlannedStartDate,
			&pm.PlannedOrderPlannedStartTime,
			&pm.PlannedOrderPlannedEndDate,
			&pm.PlannedOrderPlannedEndTime,
			&pm.LastChangeDateTime,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ProductBuyer,
			&pm.ProductSeller,
			&pm.Project,
			&pm.Reservation,
			&pm.ReservationItem,
			&pm.PlannedOrderLongText,
			&pm.PlannedOrderIsFixed,
			&pm.PlannedOrderBOMIsFixed,
			&pm.LastScheduledDate,
			&pm.ScheduledBasicEndDate,
			&pm.ScheduledBasicEndTime,
			&pm.ScheduledBasicStartDate,
			&pm.ScheduledBasicStartTime,
			&pm.SchedulingType,
			&pm.PlannedOrderIsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}
		data := pm
		res = append(res, &PlannedOrderItem{
			PlannedOrder:                          data.PlannedOrder,
			PlannedOrderItem:                      data.PlannedOrderItem,
			Product:                               data.Product,
			ProductDeliverFromParty:               data.ProductDeliverFromParty,
			ProductDeliverToParty:                 data.ProductDeliverToParty,
			IssuingPlant:                          data.IssuingPlant,
			IssuingPlantStorageLocation:           data.IssuingPlantStorageLocation,
			ReceivingPlant:                        data.ReceivingPlant,
			ReceivingPlantStorageLocation:         data.ReceivingPlantStorageLocation,
			ProductionPlantBusinessPartner:        data.ProductionPlantBusinessPartner,
			ProductionPlant:                       data.ProductionPlant,
			ProductionPlantStorageLocation:        data.ProductionPlantStorageLocation,
			BaseUnit:                              data.BaseUnit,
			MRPArea:                               data.MRPArea,
			MRPController:                         data.MRPController,
			PlannedOrderQuantityInBaseUnit:        data.PlannedOrderQuantityInBaseUnit,
			PlannedOrderPlannedScrapQtyInBaseUnit: data.PlannedOrderPlannedScrapQtyInBaseUnit,
			PlannedOrderIssuingUnit:               data.PlannedOrderIssuingUnit,
			PlannedOrderReceivingUnit:             data.PlannedOrderReceivingUnit,
			PlannedOrderIssuingQuantity:           data.PlannedOrderIssuingQuantity,
			PlannedOrderReceivingQuantity:         data.PlannedOrderReceivingQuantity,
			PlannedOrderPlannedStartDate:          data.PlannedOrderPlannedStartDate,
			PlannedOrderPlannedStartTime:          data.PlannedOrderPlannedStartTime,
			PlannedOrderPlannedEndDate:            data.PlannedOrderPlannedEndDate,
			PlannedOrderPlannedEndTime:            data.PlannedOrderPlannedEndTime,
			LastChangeDateTime:                    data.LastChangeDateTime,
			OrderID:                               data.OrderID,
			OrderItem:                             data.OrderItem,
			ProductBuyer:                          data.ProductBuyer,
			ProductSeller:                         data.ProductSeller,
			Project:                               data.Project,
			Reservation:                           data.Reservation,
			ReservationItem:                       data.ReservationItem,
			PlannedOrderLongText:                  data.PlannedOrderLongText,
			PlannedOrderIsFixed:                   data.PlannedOrderIsFixed,
			PlannedOrderBOMIsFixed:                data.PlannedOrderBOMIsFixed,
			LastScheduledDate:                     data.LastScheduledDate,
			ScheduledBasicEndDate:                 data.ScheduledBasicEndDate,
			ScheduledBasicEndTime:                 data.ScheduledBasicEndTime,
			ScheduledBasicStartDate:               data.ScheduledBasicStartDate,
			ScheduledBasicStartTime:               data.ScheduledBasicStartTime,
			SchedulingType:                        data.SchedulingType,
			PlannedOrderIsReleased:                data.PlannedOrderIsReleased,
			IsMarkedForDeletion:                   data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_planned_order_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToPlannedOrderComponentKey() *PlannedOrderComponentKey {
	pm := &requests.PlannedOrderComponentKey{
		IsMarkedForDeletion: false,
	}

	data := pm
	res := PlannedOrderComponentKey{
		PlannedOrder:        data.PlannedOrder,
		PlannedOrderItem:    data.PlannedOrderItem,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToPlannedOrderComponent(rows *sql.Rows) ([]*PlannedOrderComponent, error) {
	defer rows.Close()
	res := make([]*PlannedOrderComponent, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PlannedOrderComponent{}

		err := rows.Scan(
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.BOMItem,
			&pm.BOMItemDescription,
			&pm.BillOfMaterialCategory,
			&pm.BillOfMaterialItemNumber,
			&pm.BillOfMaterialInternalID,
			&pm.Reservation,
			&pm.ReservationItem,
			&pm.ComponentProduct,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductRequirementDate,
			&pm.ComponentProductRequiredQuantity,
			&pm.BaseUnit,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.StockConfirmationPartnerFunction,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StorageLocationForMRP,
			&pm.ComponentWithdrawnQuantity,
			&pm.ComponentScrapInPercent,
			&pm.QuantityIsFixed,
			&pm.LastChangeDateTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}
		data := pm
		res = append(res, &PlannedOrderComponent{
			PlannedOrder:                     data.PlannedOrder,
			PlannedOrderItem:                 data.PlannedOrderItem,
			BOMItem:                          data.BOMItem,
			BOMItemDescription:               data.BOMItemDescription,
			BillOfMaterialCategory:           data.BillOfMaterialCategory,
			BillOfMaterialItemNumber:         data.BillOfMaterialItemNumber,
			BillOfMaterialInternalID:         data.BillOfMaterialInternalID,
			Reservation:                      data.Reservation,
			ReservationItem:                  data.ReservationItem,
			ComponentProduct:                 data.ComponentProduct,
			ComponentProductDeliverFromParty: data.ComponentProductDeliverFromParty,
			ComponentProductDeliverToParty:   data.ComponentProductDeliverToParty,
			ComponentProductBuyer:            data.ComponentProductBuyer,
			ComponentProductSeller:           data.ComponentProductSeller,
			ComponentProductRequirementDate:  data.ComponentProductRequirementDate,
			ComponentProductRequiredQuantity: data.ComponentProductRequiredQuantity,
			BaseUnit:                         data.BaseUnit,
			MRPArea:                          data.MRPArea,
			MRPController:                    data.MRPController,
			StockConfirmationPartnerFunction: data.StockConfirmationPartnerFunction,
			StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:           data.StockConfirmationPlant,
			StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
			StorageLocationForMRP:            data.StorageLocationForMRP,
			ComponentWithdrawnQuantity:       data.ComponentWithdrawnQuantity,
			ComponentScrapInPercent:          data.ComponentScrapInPercent,
			QuantityIsFixed:                  data.QuantityIsFixed,
			LastChangeDateTime:               data.LastChangeDateTime,
			IsMarkedForDeletion:              data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_planned_order_component_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// Item
func (psdc *SDC) ConvertToProductMasterBPPlantKey() *ProductMasterBPPlantKey {
	pm := &requests.ProductMasterBPPlantKey{}

	data := pm
	res := ProductMasterBPPlantKey{
		ComponentProduct:               data.ComponentProduct,
		ProductionPlantBusinessPartner: data.ProductionPlantBusinessPartner,
		ProductionPlant:                data.ProductionPlant,
	}

	return &res
}

func (psdc *SDC) ConvertToProductMasterBPPlant(rows *sql.Rows) ([]*ProductMasterBPPlant, error) {
	defer rows.Close()
	res := make([]*ProductMasterBPPlant, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductMasterBPPlant{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.IsBatchManagementRequired,
			&pm.BatchManagementPolicy,
		)
		if err != nil {
			return nil, err
		}
		data := pm
		res = append(res, &ProductMasterBPPlant{
			Product:                   data.Product,
			BusinessPartner:           data.BusinessPartner,
			Plant:                     data.Plant,
			IsBatchManagementRequired: data.IsBatchManagementRequired,
			BatchManagementPolicy:     data.BatchManagementPolicy,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_bp_plant_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToCreationDateHeader(systemDate string) *CreationDate {
	pm := &requests.CreationDate{}

	pm.CreationDate = systemDate

	data := pm
	res := CreationDate{
		CreationDate: data.CreationDate,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeDateHeader(systemDate string) *LastChangeDate {
	pm := &requests.LastChangeDate{}

	pm.LastChangeDate = systemDate

	data := pm
	res := LastChangeDate{
		LastChangeDate: data.LastChangeDate,
	}

	return &res
}
