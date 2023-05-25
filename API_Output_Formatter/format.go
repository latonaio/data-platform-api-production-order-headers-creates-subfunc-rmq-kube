package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*Header, error) {
	var err error

	header := &Header{}
	inputHeader := sdc.Header

	// 入力ファイル
	header, err = jsonTypeConversion(header, inputHeader)
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	header.ProductionOrderType = getStringPtr("PP")
	header.CreationDate = psdc.CreationDateHeader.CreationDate
	header.LastChangeDate = psdc.LastChangeDateHeader.LastChangeDate
	header.HeaderIsReleased = sdc.InputParameters.ReleaseProductionOrder
	header.HeaderIsPartiallyConfirmed = psdc.HeaderIsPartiallyConfirmed.HeaderIsPartiallyConfirmed
	header.HeaderIsConfirmed = psdc.HeaderIsConfirmed.HeaderIsConfirmed
	header.HeaderIsLocked = getBoolPtr(false)
	header.HeaderIsMarkedForDeletion = getBoolPtr(false)
	header.Product = psdc.PlannedOrderHeader[0].Product
	header.OwnerProductionPlant = getValue(psdc.PlannedOrderHeader[0].OwnerProductionPlant)
	header.OwnerProductionPlantBusinessPartner = getValue(psdc.PlannedOrderHeader[0].OwnerProductionPlantBusinessPartner)
	header.OwnerProductionPlantStorageLocation = psdc.PlannedOrderHeader[0].OwnerProductionPlantStorageLocation
	header.MRPArea = psdc.PlannedOrderHeader[0].MRPArea
	header.MRPController = psdc.PlannedOrderHeader[0].MRPController
	// header.ProductionVersion = getStringPtr("0001")  //ProductionVersionがstring→intになった
	header.PlannedOrder = &psdc.PlannedOrderHeader[0].PlannedOrder
	header.OrderID = psdc.PlannedOrderHeader[0].OrderID
	header.OrderItem = psdc.PlannedOrderHeader[0].OrderItem
	header.ProductionOrderPlannedStartDate = psdc.PlannedOrderHeader[0].PlannedOrderPlannedStartDate
	header.ProductionOrderPlannedStartTime = psdc.PlannedOrderHeader[0].PlannedOrderPlannedStartTime
	header.ProductionOrderPlannedEndDate = psdc.PlannedOrderHeader[0].PlannedOrderPlannedEndDate
	header.ProductionOrderPlannedEndTime = psdc.PlannedOrderHeader[0].PlannedOrderPlannedEndTime
	header.ProductionUnit = psdc.PlannedOrderHeader[0].PlannedOrderOriginIssuingUnit
	header.TotalQuantity = getValue(psdc.TotalQuantityHeader.TotalQuantity)
	header.PlannedScrapQuantity = psdc.PlannedScrapQuantityHeader[0].PlannedScrapQuantity
	header.ConfirmedYieldQuantity = psdc.TotalQuantityHeader.TotalQuantity

	return header, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func getValue[T any](ptr *T) T {
	var zero T
	if ptr == nil {
		return zero
	}
	return *ptr
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
