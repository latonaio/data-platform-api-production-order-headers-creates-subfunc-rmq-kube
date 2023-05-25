package subfunction

import (
	"context"
	api_input_reader "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-production-order-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-production-order-headers-creates-subfunc-rmq-kube/config"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type SubFunction struct {
	ctx  context.Context
	db   *database.Mysql
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	l    *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, conf *config.Conf, rmq *rabbitmq.RabbitmqClient, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx:  ctx,
		db:   db,
		conf: conf,
		rmq:  rmq,
		l:    l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) ProcessType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.ProcessType, error) {
	processType := psdc.ConvertToProcessType()

	if isBulkProcess(sdc, processType) {
		processType.BulkProcess = true
	}

	// if isIndividualProcess(sdc, processType, referenceType) {
	// 	processType.IndividualProcess = true
	// }

	if !processType.BulkProcess && !processType.IndividualProcess {
		return nil, xerrors.Errorf("一括登録または個別登録に必要な入力パラメータが揃っていません。")
	}

	return processType, nil
}

func isBulkProcess(
	sdc *api_input_reader.SDC,
	processType *api_processing_data_formatter.ProcessType,
) bool {
	inputParameters := sdc.InputParameters

	if inputParameters.MRPArea != nil && inputParameters.OwnerProductionPlantBusinessPartner != nil &&
		inputParameters.OwnerProductionPlant != nil && inputParameters.ProductInHeader != nil &&
		inputParameters.ProductInItem != nil && inputParameters.ProductionPlantBusinessPartner != nil &&
		inputParameters.ProductionPlant != nil && inputParameters.ProductionPlantStorageLocation != nil {
		if (*inputParameters.MRPArea)[0] != nil && (*inputParameters.OwnerProductionPlantBusinessPartner)[0] != nil &&
			(*inputParameters.OwnerProductionPlant)[0] != nil && (*inputParameters.ProductInHeader)[0] != nil &&
			(*inputParameters.ProductInItem)[0] != nil && (*inputParameters.ProductionPlantBusinessPartner)[0] != nil &&
			(*inputParameters.ProductionPlant)[0] != nil && (*inputParameters.ProductionPlantStorageLocation)[0] != nil {
			if len(*(*inputParameters.MRPArea)[0]) != 0 && len(*(*inputParameters.OwnerProductionPlant)[0]) != 0 &&
				len(*(*inputParameters.ProductInHeader)[0]) != 0 && len(*(*inputParameters.ProductInItem)[0]) != 0 &&
				len(*(*inputParameters.ProductionPlant)[0]) != 0 && len(*(*inputParameters.ProductionPlantStorageLocation)[0]) != 0 {
				processType.PluralitySpec = true
				return true
			}
		}
	}
	if inputParameters.MRPAreaTo != nil && inputParameters.MRPAreaFrom != nil &&
		inputParameters.OwnerProductionPlantBusinessPartnerTo != nil && inputParameters.OwnerProductionPlantBusinessPartnerFrom != nil &&
		inputParameters.OwnerProductionPlantTo != nil && inputParameters.OwnerProductionPlantFrom != nil &&
		inputParameters.ProductInHeaderTo != nil && inputParameters.ProductInHeaderFrom != nil &&
		inputParameters.ProductInItemTo != nil && inputParameters.ProductInItemFrom != nil &&
		inputParameters.ProductionPlantBusinessPartnerTo != nil && inputParameters.ProductionPlantBusinessPartnerFrom != nil &&
		inputParameters.ProductionPlantTo != nil && inputParameters.ProductionPlantFrom != nil &&
		inputParameters.ProductionPlantStorageLocationTo != nil && inputParameters.ProductionPlantStorageLocationFrom != nil {
		if len(*inputParameters.MRPAreaTo) != 0 && len(*inputParameters.MRPAreaFrom) != 0 &&
			len(*inputParameters.OwnerProductionPlantTo) != 0 && len(*inputParameters.OwnerProductionPlantFrom) != 0 &&
			len(*inputParameters.ProductInHeaderTo) != 0 && len(*inputParameters.ProductInHeaderFrom) != 0 &&
			len(*inputParameters.ProductInItemTo) != 0 && len(*inputParameters.ProductInItemFrom) != 0 &&
			len(*inputParameters.ProductionPlantTo) != 0 && len(*inputParameters.ProductionPlantFrom) != 0 &&
			len(*inputParameters.ProductionPlantStorageLocationTo) != 0 && len(*inputParameters.ProductionPlantStorageLocationFrom) != 0 {
			processType.RangeSpec = true
			return true
		}
	}

	return false
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	psdc.MetaData = f.MetaData(sdc, psdc)
	psdc.ProcessType, err = f.ProcessType(sdc, psdc)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1-1.計画手配ヘッダデータの取得
		psdc.PlannedOrderHeader, e = f.PlannedOrderHeaderInBulkProcess(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-1-2.計画手配明細データの取得  //1-1-1
		psdc.PlannedOrderItem, e = f.PlannedOrderItemInBulkProcess(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-1-3.計画手配構成品目データの取得  //1-1-2
		psdc.PlannedOrderComponent, e = f.PlannedOrderComponent(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-1-0. 在庫確認実行の判断
		psdc.ExecuteProductAvailabilityCheck, e = f.ExecuteProductAvailabilityCheck(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		if psdc.ExecuteProductAvailabilityCheck.ExecuteProductAvailabilityCheck {
			// 2-1-1. ロット管理品目の確認(品目マスタBPプラントデータの取得)  //1-1-2,1-1-3
			psdc.ProductMasterBPPlant, e = f.ProductMasterBPPlant(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-3-1. ロットマスタデータの取得 //1-1-3
			psdc.BatchMasterRecordBatch, e = f.BatchMasterRecordBatch(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-2-1. 構成品目ごとの在庫確認①(通常の在庫確認)   //1-1-2,1-1-3,2-1-1,2-3-1
			psdc.StockConfirmation, e = f.StockConfirmation(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}

		// 2-4-0. ItemIsPartiallyConfirmedのセット  //2-2-1
		psdc.ItemIsPartiallyConfirmed = f.ItemIsPartiallyConfirmed(sdc, psdc)

		// 2-4-1. ItemIsConfirmedのセット  //2-2-1
		psdc.ItemIsConfirmed = f.ItemIsConfirmed(sdc, psdc)

		// 2-6. 製造指図明細への数量のセット  //2-2-1
		psdc.TotalQuantity = f.TotalQuantity(sdc, psdc)

		// 11-1-1. HeaderIsPartiallyConfirmedのセット
		psdc.HeaderIsPartiallyConfirmed = f.HeaderIsPartiallyConfirmed(sdc, psdc)

		// 11-1-2. HeaderIsConfirmedのセット
		psdc.HeaderIsConfirmed = f.HeaderIsConfirmed(sdc, psdc)

		// 11-2-1. TotalQuantityのセット  //2-6
		psdc.TotalQuantityHeader = f.TotalQuantityHeader(sdc, psdc)

		// 1-4. PlannedScrapQuantity	//1-1-3,11-2-1
		psdc.PlannedScrapQuantityHeader = f.PlannedScrapQuantityHeader(sdc, psdc)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 99-1-1. CreationDate(Header)
		psdc.CreationDateHeader = f.CreationDateHeader(sdc, psdc)

		// 99-2-1. LastChangeDate(Header)
		psdc.LastChangeDateHeader = f.LastChangeDateHeader(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)
	err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
