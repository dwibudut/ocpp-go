package tariffcost

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// -------------------- Cost Updated (CSMS -> CS) --------------------

const SetDefaultTariffFeatureName = "SetDefaultTariff"

type TariffSetStatusEnumType string

const (
	TariffSetStatusAccepted              TariffSetStatusEnumType = "Accepted"
	TariffSetStatusRejected              TariffSetStatusEnumType = "Rejected"
	TariffSetStatusTooManyElements       TariffSetStatusEnumType = "TooManyElements"
	TariffSetStatusConditionNotSupported TariffSetStatusEnumType = "ConditionNotSupported"
	TariffSetStatusDuplicateTariffId     TariffSetStatusEnumType = "DuplicateTariffId"
)

func isValidTariffSetStatus(fl validator.FieldLevel) bool {
	status := TariffSetStatusEnumType(fl.Field().String())
	switch status {
	case TariffSetStatusAccepted, TariffSetStatusRejected, TariffSetStatusTooManyElements, TariffSetStatusConditionNotSupported, TariffSetStatusDuplicateTariffId:
		return true
	default:
		return false
	}
}

// The field definition of the SetDefaultTariff request payload sent by the CSMS to the Charging Station.
type SetDefaultTariffRequest struct {
	EvseId int              `json:"evseId" validate:"gte=0"`
	Tariff types.TariffType `json:"tariff" validate:"required"`
}

// This field definition of the SetDefaultTariff response payload, sent by the Charging Station to the CSMS in response to a SetDefaultTariffRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type SetDefaultTariffResponse struct {
	Status     TariffSetStatusEnumType `json:"status" validate:"required,tariffSetStatus"`
	StatusInfo *types.StatusInfo       `json:"statusInfo,omitempty" validate:"omitempty"`
}

// The driver wants to know how much the running total cost is, updated at a relevant interval, while a transaction is ongoing.
// To fulfill this requirement, the CSMS sends a SetDefaultTariffRequest to the Charging Station to update the current total cost, every Y seconds.
// Upon receipt of the SetDefaultTariffRequest, the Charging Station responds with a SetDefaultTariffResponse, then shows the updated cost to the driver.
type SetDefaultTariffFeature struct{}

func (f SetDefaultTariffFeature) GetFeatureName() string {
	return SetDefaultTariffFeatureName
}

func (f SetDefaultTariffFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(SetDefaultTariffRequest{})
}

func (f SetDefaultTariffFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(SetDefaultTariffResponse{})
}

func (r SetDefaultTariffRequest) GetFeatureName() string {
	return SetDefaultTariffFeatureName
}

func (c SetDefaultTariffResponse) GetFeatureName() string {
	return SetDefaultTariffFeatureName
}

// Creates a new SetDefaultTariffRequest, containing all required fields. There are no optional fields for this message.
func NewSetDefaultTariffRequest(evseId int, tariff types.TariffType) *SetDefaultTariffRequest {
	return &SetDefaultTariffRequest{EvseId: evseId, Tariff: tariff}
}

// Creates a new SetDefaultTariffResponse, which doesn't contain any required or optional fields.
func NewSetDefaultTariffResponse(status TariffSetStatusEnumType) *SetDefaultTariffResponse {
	return &SetDefaultTariffResponse{Status: status}
}

func init() {
	_ = types.Validate.RegisterValidation("tariffSetStatus", isValidTariffSetStatus)
}
