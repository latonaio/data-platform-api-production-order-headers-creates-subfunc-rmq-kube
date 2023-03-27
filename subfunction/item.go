package subfunction

import (
	api_input_reader "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) ProductMasterBPPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterBPPlant, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductMasterBPPlantKey()

	for _, component := range psdc.PlannedOrderComponent {
		if component.ComponentProduct == nil {
			continue
		}
		dataKey.ComponentProduct = append(dataKey.ComponentProduct, *component.ComponentProduct)

		plannedOrder := component.PlannedOrder
		plannedOrderItem := component.PlannedOrderItem

		for _, item := range psdc.PlannedOrderItem {
			if plannedOrder == item.PlannedOrder && plannedOrderItem == item.PlannedOrderItem {
				if item.ProductionPlantBusinessPartner == nil || item.ProductionPlant == nil {
					continue
				}
				dataKey.ProductionPlantBusinessPartner = append(dataKey.ProductionPlantBusinessPartner, *item.ProductionPlantBusinessPartner)
				dataKey.ProductionPlant = append(dataKey.ProductionPlant, *item.ProductionPlant)
				break
			}
		}
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey.ComponentProduct)-1) + "(?,?,?)"
	for i := range dataKey.ComponentProduct {
		args = append(args, dataKey.ComponentProduct[i], dataKey.ProductionPlantBusinessPartner[i], dataKey.ProductionPlant[i])
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, IsBatchManagementRequired, BatchManagementPolicy
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductMasterBPPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
