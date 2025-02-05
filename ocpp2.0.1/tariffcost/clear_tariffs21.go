package tariffcost

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// -------------------- Cost Updated (CSMS -> CS) --------------------

const ClearTariffsFeatureName = "ClearTariffs"

type TariffClearStatusEnumType string

const (
	TariffClearStatusAccepted TariffClearStatusEnumType = "Accepted"
	TariffClearStatusRejected TariffClearStatusEnumType = "Rejected"
	TariffClearStatusNoTariff TariffClearStatusEnumType = "NoTariff"
)

func isValidTariffClearStatus(fl validator.FieldLevel) bool {
	status := TariffClearStatusEnumType(fl.Field().String())
	switch status {
	case TariffClearStatusAccepted, TariffClearStatusRejected, TariffClearStatusNoTariff:
		return true
	default:
		return false
	}
}

// ClearTariffsResultType
type ClearTariffsResultType struct {
	StatusInfo *types.StatusInfo       `json:"statusInfo,omitempty" validate:"omitempty"`
	Status     TariffSetStatusEnumType `json:"status" validate:"required,tariffClearStatus"`
	TariffId   string                  `json:"tariffId,omitempty" validate:"omitempty,max=60"`
}

// The field definition of the ClearTariffs request payload sent by the CSMS to the Charging Station.
type ClearTariffsRequest struct {
	TariffIds []string `json:"tariffIds,omitempty" validate:"omitempty,min=1,dive,max=60"`
	EvseId    int      `json:"evseId" validate:"gte=0"`
}

// This field definition of the ClearTariffs response payload, sent by the Charging Station to the CSMS in response to a ClearTariffsRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type ClearTariffsResponse struct {
	ClearTariffsResult []ClearTariffsResultType `json:"clearTariffsResult" validate:"required,min=1,dive"`
}

// The driver wants to know how much the running total cost is, updated at a relevant interval, while a transaction is ongoing.
// To fulfill this requirement, the CSMS sends a ClearTariffsRequest to the Charging Station to update the current total cost, every Y seconds.
// Upon receipt of the ClearTariffsRequest, the Charging Station responds with a ClearTariffsResponse, then shows the updated cost to the driver.
type ClearTariffsFeature struct{}

func (f ClearTariffsFeature) GetFeatureName() string {
	return ClearTariffsFeatureName
}

func (f ClearTariffsFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(ClearTariffsRequest{})
}

func (f ClearTariffsFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(ClearTariffsResponse{})
}

func (r ClearTariffsRequest) GetFeatureName() string {
	return ClearTariffsFeatureName
}

func (c ClearTariffsResponse) GetFeatureName() string {
	return ClearTariffsFeatureName
}

// Creates a new ClearTariffsRequest, containing all required fields. There are no optional fields for this message.
func NewClearTariffsRequest() *ClearTariffsRequest {
	return &ClearTariffsRequest{}
}

// Creates a new ClearTariffsResponse, which doesn't contain any required or optional fields.
func NewClearTariffsResponse(clearTariffsResult []ClearTariffsResultType) *ClearTariffsResponse {
	return &ClearTariffsResponse{ClearTariffsResult: clearTariffsResult}
}

func init() {
	_ = types.Validate.RegisterValidation("tariffClearStatus", isValidTariffClearStatus)
}
