package ocpp2

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/diagnostics"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/iso15118"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/tariffcost"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/transactions"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

func (cs *chargingStation) ClosePeriodicEventStream(id int, props ...func(request *diagnostics.ClosePeriodicEventStreamRequest)) (*diagnostics.ClosePeriodicEventStreamResponse, error) {
	request := diagnostics.NewClosePeriodicEventStreamRequest(id)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*diagnostics.ClosePeriodicEventStreamResponse), err
	}
}

func (cs *chargingStation) NotifyPeriodicEventStream(id, pending int, basetime *types.DateTime, data []diagnostics.StreamDataElementType, props ...func(request *diagnostics.NotifyPeriodicEventStreamRequest)) error {
	request := diagnostics.NewNotifyPeriodicEventStreamRequest(id, pending, basetime, data)
	for _, fn := range props {
		fn(request)
	}
	err := cs.SendRequestAsync(request, nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (cs *chargingStation) NotifyPriorityCharging(transactionId string, activated bool, props ...func(request *diagnostics.NotifyPriorityChargingRequest)) (*diagnostics.NotifyPriorityChargingResponse, error) {
	request := diagnostics.NewNotifyPriorityChargingRequest(transactionId, activated)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*diagnostics.NotifyPriorityChargingResponse), err
	}
}

func (cs *chargingStation) NotifySettlement(pspRef string, status types.PaymentStatusEnumType, settlementAmount float64, settlementTime *types.DateTime, props ...func(request *diagnostics.NotifySettlementRequest)) (*diagnostics.NotifySettlementResponse, error) {
	request := diagnostics.NewNotifySettlementRequest(pspRef, status, settlementAmount, settlementTime)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*diagnostics.NotifySettlementResponse), err
	}
}

func (cs *chargingStation) OpenPeriodicEventStream(constantStreamData diagnostics.ConstantStreamDataType, props ...func(request *diagnostics.OpenPeriodicEventStreamRequest)) (*diagnostics.OpenPeriodicEventStreamResponse, error) {
	request := diagnostics.NewOpenPeriodicEventStreamRequest(constantStreamData)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*diagnostics.OpenPeriodicEventStreamResponse), err
	}
}

func (cs *chargingStation) NotifyDERAlarm(controlType types.DERControlEnumType, timestamp *types.DateTime, props ...func(request *smartcharging.NotifyDERAlarmRequest)) (*smartcharging.NotifyDERAlarmResponse, error) {
	request := smartcharging.NewNotifyDERAlarmRequest(controlType, timestamp)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*smartcharging.NotifyDERAlarmResponse), err
	}
}

func (cs *chargingStation) NotifyDERStartStop(controlId string, started bool, timestamp *types.DateTime, props ...func(request *smartcharging.NotifyDERStartStopRequest)) (*smartcharging.NotifyDERStartStopResponse, error) {
	request := smartcharging.NewNotifyDERStartStopRequest(controlId, started, timestamp)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*smartcharging.NotifyDERStartStopResponse), err
	}
}

func (cs *chargingStation) PullDynamicScheduleUpdate(chargingProfileId int, props ...func(request *smartcharging.PullDynamicScheduleUpdateRequest)) (*smartcharging.PullDynamicScheduleUpdateResponse, error) {
	request := smartcharging.NewPullDynamicScheduleUpdateRequest(chargingProfileId)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*smartcharging.PullDynamicScheduleUpdateResponse), err
	}
}

func (cs *chargingStation) ReportDERControl(requestId int, props ...func(request *smartcharging.ReportDERControlRequest)) (*smartcharging.ReportDERControlResponse, error) {
	request := smartcharging.NewReportDERControlRequest(requestId)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*smartcharging.ReportDERControlResponse), err
	}
}

func (cs *chargingStation) VatNumberValidation(vatNumber string, props ...func(request *tariffcost.VatNumberValidationRequest)) (*tariffcost.VatNumberValidationResponse, error) {
	request := tariffcost.NewVatNumberValidationRequest(vatNumber)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*tariffcost.VatNumberValidationResponse), err
	}
}

func (cs *chargingStation) BatterySwap(batteryData []transactions.BatteryDataType, eventType transactions.BatterySwapEventEnumType, idToken *types.IdToken, requestId int, props ...func(request *transactions.BatterySwapRequest)) (*transactions.BatterySwapResponse, error) {
	request := transactions.NewBatterySwapRequest(batteryData, eventType, idToken, requestId)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*transactions.BatterySwapResponse), err
	}
}

func (cs *chargingStation) GetCertificateChainStatus(certificateStatusRequests []iso15118.CertificateStatusRequestInfoType, props ...func(request *iso15118.GetCertificateChainStatusRequest)) (*iso15118.GetCertificateChainStatusResponse, error) {
	request := iso15118.NewGetCertificateChainStatusRequest(certificateStatusRequests)
	for _, fn := range props {
		fn(request)
	}
	response, err := cs.SendRequest(request)
	if err != nil {
		return nil, err
	} else {
		return response.(*iso15118.GetCertificateChainStatusResponse), err
	}
}
