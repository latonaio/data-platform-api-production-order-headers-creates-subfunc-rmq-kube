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

	// I-1-1
	header, err = jsonTypeConversion(header, psdc.PlannedOrderHeder[0])
	if err != nil {
		return nil, xerrors.Errorf("request create error: %w", err)
	}

	header.ProductionOrderType = getStringPtr("PP")
	header.CreationDate = &psdc.CreationDateHeader.CreationDate
	header.LastChangeDate = &psdc.LastChangeDateHeader.LastChangeDate
	// header.HeaderIsPartiallyConfirmed =
	// header.HeaderIsConfirmed =
	header.HeaderIsLocked = getBoolPtr(false)
	header.HeaderIsMarkedForDeletion = getBoolPtr(false)
	header.ProductionVersion = getStringPtr("0001")
	// header.TotalQuantity =
	// header.PlannedScrapQuantity =
	// header.ConfirmedYieldQuantity =

	return header, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
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
