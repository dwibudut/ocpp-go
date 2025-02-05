// The authorization functional block contains OCPP 2.0 features that show tariff and costs to an EV driver, when supported by the charging station.
package tariffcost

import "github.com/lorenzodonini/ocpp-go/ocpp"

// Needs to be implemented by a CSMS for handling messages part of the OCPP 2.0 Tariff and cost profile.
type CSMSHandler interface {
	// OnVatNumberValidation is called on the CSMS whenever a VatNumberValidationRequest is received from a charging station.
	OnVatNumberValidation(chargingStationID string, request *VatNumberValidationRequest) (response *VatNumberValidationResponse, err error)
}

// Needs to be implemented by Charging stations for handling messages part of the OCPP 2.0 Tariff and cost profile.
type ChargingStationHandler interface {
	// OnCostUpdated is called on a charging station whenever a CostUpdatedRequest is received from the CSMS.
	OnCostUpdated(request *CostUpdatedRequest) (confirmation *CostUpdatedResponse, err error)
	// OnChangeTransactionTariff is called on a charging station whenever a ChangeTransactionTariffRequest is received from the CSMS.
	OnChangeTransactionTariff(request *ChangeTransactionTariffRequest) (confirmation *ChangeTransactionTariffResponse, err error)
	// OnClearTariffs is called on a charging station whenever a ClearTariffsRequest is received from the CSMS.
	OnClearTariffs(request *ClearTariffsRequest) (confirmation *ClearTariffsResponse, err error)
	// OnGetTariffs is called on a charging station whenever a GetTariffsRequest is received from the CSMS.
	OnGetTariffs(request *GetTariffsRequest) (confirmation *GetTariffsResponse, err error)
	// OnSetDefaultTariff is called on a charging station whenever a SetDefaultTariffRequest is received from the CSMS.
	OnSetDefaultTariff(request *SetDefaultTariffRequest) (confirmation *SetDefaultTariffResponse, err error)
}

const ProfileName = "tariffCost"

var Profile = ocpp.NewProfile(
	ProfileName,
	CostUpdatedFeature{},
	ChangeTransactionTariffFeature{},
	ClearTariffsFeature{},
	GetTariffsFeature{},
	SetDefaultTariffFeature{},
	VatNumberValidationFeature{},
)
