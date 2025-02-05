package ocpp2_test

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/diagnostics"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/iso15118"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/remotecontrol"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/tariffcost"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/transactions"
)

func (websocket MockWebSocket) OcppVersion() string {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CS DIAGNOSTICS HANDLER ----------------------
func (handler *MockChargingStationDiagnosticsHandler) OnAdjustPeriodicEventStream(request *diagnostics.AdjustPeriodicEventStreamRequest) (response *diagnostics.AdjustPeriodicEventStreamResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationDiagnosticsHandler) OnGetPeriodicEventStream(request *diagnostics.GetPeriodicEventStreamRequest) (response *diagnostics.GetPeriodicEventStreamResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationDiagnosticsHandler) OnNotifyWebPaymentStarted(request *diagnostics.NotifyWebPaymentStartedRequest) (response *diagnostics.NotifyWebPaymentStartedResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CSMS DIAGNOSTICS HANDLER ----------------------

func (handler *MockCSMSDiagnosticsHandler) OnClosePeriodicEventStream(chargingStationID string, request *diagnostics.ClosePeriodicEventStreamRequest) (response *diagnostics.ClosePeriodicEventStreamResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockCSMSDiagnosticsHandler) OnNotifyPriorityCharging(chargingStationID string, request *diagnostics.NotifyPriorityChargingRequest) (response *diagnostics.NotifyPriorityChargingResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockCSMSDiagnosticsHandler) OnNotifySettlement(chargingStationID string, request *diagnostics.NotifySettlementRequest) (response *diagnostics.NotifySettlementResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockCSMSDiagnosticsHandler) OnOpenPeriodicEventStream(chargingStationID string, request *diagnostics.OpenPeriodicEventStreamRequest) (response *diagnostics.OpenPeriodicEventStreamResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockCSMSDiagnosticsHandler) OnNotifyPeriodicEventStream(chargingStationID string, request *diagnostics.NotifyPeriodicEventStreamRequest) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CSMS ISO15118 HANDLER ----------------------

func (handler *MockCSMSIso15118Handler) OnGetCertificateChainStatus(chargingStationID string, request *iso15118.GetCertificateChainStatusRequest) (response *iso15118.GetCertificateChainStatusResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CS REMOTE CONTROL HANDLER ----------------------

func (handler *MockChargingStationRemoteControlHandler) OnRequestBatterySwap(request *remotecontrol.RequestBatterySwapRequest) (response *remotecontrol.RequestBatterySwapResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CS SMART CHARGING HANDLER ----------------------

func (handler *MockChargingStationSmartChargingHandler) OnAFRRSignal(request *smartcharging.AFRRSignalRequest) (response *smartcharging.AFRRSignalResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationSmartChargingHandler) OnClearDERControl(request *smartcharging.ClearDERControlRequest) (response *smartcharging.ClearDERControlResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationSmartChargingHandler) OnNotifyAllowedEnergyTransfer(request *smartcharging.NotifyAllowedEnergyTransferRequest) (response *smartcharging.NotifyAllowedEnergyTransferResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationSmartChargingHandler) OnGetDERControl(request *smartcharging.GetDERControlRequest) (response *smartcharging.GetDERControlResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationSmartChargingHandler) OnSetDERControl(request *smartcharging.SetDERControlRequest) (response *smartcharging.SetDERControlResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationSmartChargingHandler) OnUpdateDynamicSchedule(request *smartcharging.UpdateDynamicScheduleRequest) (response *smartcharging.UpdateDynamicScheduleResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationSmartChargingHandler) OnUsePriorityCharging(request *smartcharging.UsePriorityChargingRequest) (response *smartcharging.UsePriorityChargingResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CSMS SMART CHARGING HANDLER ----------------------

func (handler *MockCSMSSmartChargingHandler) OnPullDynamicScheduleUpdate(chargingStationID string, request *smartcharging.PullDynamicScheduleUpdateRequest) (reponse *smartcharging.PullDynamicScheduleUpdateResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockCSMSSmartChargingHandler) OnNotifyDERAlarm(chargingStationID string, request *smartcharging.NotifyDERAlarmRequest) (reponse *smartcharging.NotifyDERAlarmResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockCSMSSmartChargingHandler) OnNotifyDERStartStop(chargingStationID string, request *smartcharging.NotifyDERStartStopRequest) (reponse *smartcharging.NotifyDERStartStopResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockCSMSSmartChargingHandler) OnReportDERControl(chargingStationID string, request *smartcharging.ReportDERControlRequest) (reponse *smartcharging.ReportDERControlResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CS TARIFF COST HANDLER ----------------------

func (handler *MockChargingStationTariffCostHandler) OnChangeTransactionTariff(request *tariffcost.ChangeTransactionTariffRequest) (confirmation *tariffcost.ChangeTransactionTariffResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationTariffCostHandler) OnClearTariffs(request *tariffcost.ClearTariffsRequest) (confirmation *tariffcost.ClearTariffsResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationTariffCostHandler) OnGetTariffs(request *tariffcost.GetTariffsRequest) (confirmation *tariffcost.GetTariffsResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (handler *MockChargingStationTariffCostHandler) OnSetDefaultTariff(request *tariffcost.SetDefaultTariffRequest) (confirmation *tariffcost.SetDefaultTariffResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CSMS TARIFF COST HANDLER ----------------------

func (m *MockCSMSTariffCostHandler) OnVatNumberValidation(chargingStationID string, request *tariffcost.VatNumberValidationRequest) (response *tariffcost.VatNumberValidationResponse, err error) {
	//TODO implement me
	panic("implement me")
}

// ---------------------- MOCK CSMS TRANSACTIONS HANDLER ----------------------

func (handler *MockCSMSTransactionsHandler) OnBatterySwap(chargingStationID string, request *transactions.BatterySwapRequest) (response *transactions.BatterySwapResponse, err error) {
	//TODO implement me
	panic("implement me")
}
