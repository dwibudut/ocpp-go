package ocpp2

import (
	"github.com/lorenzodonini/ocpp-go/ocpp"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/diagnostics"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/remotecontrol"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/tariffcost"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

func (cs *csms) UpdateDynamicSchedule(clientId string, callback func(*smartcharging.UpdateDynamicScheduleResponse, error), chargingProfileId int, scheduleUpdate types.ChargingScheduleUpdateType, props ...func(request *smartcharging.UpdateDynamicScheduleRequest)) error {
	request := smartcharging.NewUpdateDynamicScheduleRequest(chargingProfileId, scheduleUpdate)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*smartcharging.UpdateDynamicScheduleResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) UsePriorityCharging(clientId string, callback func(*smartcharging.UsePriorityChargingResponse, error), transactionId string, activate bool, props ...func(request *smartcharging.UsePriorityChargingRequest)) error {
	request := smartcharging.NewUsePriorityChargingRequest(transactionId, activate)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*smartcharging.UsePriorityChargingResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) AdjustPeriodicEventStream(clientId string, callback func(*diagnostics.AdjustPeriodicEventStreamResponse, error), id int, params diagnostics.PeriodicEventStreamParamsType, props ...func(request *diagnostics.AdjustPeriodicEventStreamRequest)) error {
	request := diagnostics.NewAdjustPeriodicEventStreamRequest(id, params)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*diagnostics.AdjustPeriodicEventStreamResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) GetPeriodicEventStream(clientId string, callback func(*diagnostics.GetPeriodicEventStreamResponse, error), props ...func(request *diagnostics.GetPeriodicEventStreamRequest)) error {
	request := diagnostics.NewGetPeriodicEventStreamRequest()
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*diagnostics.GetPeriodicEventStreamResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) NotifyWebPaymentStarted(clientId string, callback func(*diagnostics.NotifyWebPaymentStartedResponse, error), evseId int, timeout int, props ...func(request *diagnostics.NotifyWebPaymentStartedRequest)) error {
	request := diagnostics.NewNotifyWebPaymentStartedRequest(evseId, timeout)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*diagnostics.NotifyWebPaymentStartedResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) RequestBatterySwap(clientId string, callback func(*remotecontrol.RequestBatterySwapResponse, error), requestId int, idToken types.IdToken, props ...func(request *remotecontrol.RequestBatterySwapRequest)) error {
	request := remotecontrol.NewRequestBatterySwapRequest(requestId, idToken)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*remotecontrol.RequestBatterySwapResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) AFRRSignal(clientId string, callback func(*smartcharging.AFRRSignalResponse, error), timestamp *types.DateTime, signal int, props ...func(request *smartcharging.AFRRSignalRequest)) error {
	request := smartcharging.NewAFRRSignalRequest(timestamp, signal)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*smartcharging.AFRRSignalResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) ClearDERControl(clientId string, callback func(*smartcharging.ClearDERControlResponse, error), isDefault bool, props ...func(request *smartcharging.ClearDERControlRequest)) error {
	request := smartcharging.NewClearDERControlRequest(isDefault)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*smartcharging.ClearDERControlResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) NotifyAllowedEnergyTransfer(clientId string, callback func(*smartcharging.NotifyAllowedEnergyTransferResponse, error), transactionId string, allowedEnergyTransfer []types.EnergyTransferModeEnumType, props ...func(request *smartcharging.NotifyAllowedEnergyTransferRequest)) error {
	request := smartcharging.NewNotifyAllowedEnergyTransferRequest(transactionId, allowedEnergyTransfer)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*smartcharging.NotifyAllowedEnergyTransferResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) GetDERControl(clientId string, callback func(*smartcharging.GetDERControlResponse, error), requestId int, props ...func(request *smartcharging.GetDERControlRequest)) error {
	request := smartcharging.NewGetDERControlRequest(requestId)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*smartcharging.GetDERControlResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) SetDERControl(clientId string, callback func(*smartcharging.SetDERControlResponse, error), isDefault bool, controlId string, controlType types.DERControlEnumType, props ...func(request *smartcharging.SetDERControlRequest)) error {
	request := smartcharging.NewSetDERControlRequest(isDefault, controlId, controlType)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*smartcharging.SetDERControlResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) ChangeTransactionTariff(clientId string, callback func(*tariffcost.ChangeTransactionTariffResponse, error), tariff types.TariffType, transactionID string, props ...func(request *tariffcost.ChangeTransactionTariffRequest)) error {
	request := tariffcost.NewChangeTransactionTariffRequest(tariff, transactionID)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*tariffcost.ChangeTransactionTariffResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) ClearTariffs(clientId string, callback func(*tariffcost.ClearTariffsResponse, error), props ...func(request *tariffcost.ClearTariffsRequest)) error {
	request := tariffcost.NewClearTariffsRequest()
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*tariffcost.ClearTariffsResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) GetTariffs(clientId string, callback func(*tariffcost.GetTariffsResponse, error), evseId int, props ...func(request *tariffcost.GetTariffsRequest)) error {
	request := tariffcost.NewGetTariffsRequest(evseId)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*tariffcost.GetTariffsResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}

func (cs *csms) SetDefaultTariff(clientId string, callback func(*tariffcost.SetDefaultTariffResponse, error), evseId int, tariff types.TariffType, props ...func(request *tariffcost.SetDefaultTariffRequest)) error {
	request := tariffcost.NewSetDefaultTariffRequest(evseId, tariff)
	for _, fn := range props {
		fn(request)
	}
	genericCallback := func(response ocpp.Response, protoError error) {
		if response != nil {
			callback(response.(*tariffcost.SetDefaultTariffResponse), protoError)
		} else {
			callback(nil, protoError)
		}
	}
	return cs.SendRequestAsync(clientId, request, genericCallback)
}
