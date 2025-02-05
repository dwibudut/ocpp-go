package tariffcost

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// -------------------- Cost Updated (CSMS -> CS) --------------------

const ChangeTransactionTariffFeatureName = "ChangeTransactionTariff"

type TariffChangeStatusEnumType string

const (
	TariffChangeStatusAccepted              TariffChangeStatusEnumType = "Accepted"
	TariffChangeStatusRejected              TariffChangeStatusEnumType = "Rejected"
	TariffChangeStatusTooManyElements       TariffChangeStatusEnumType = "TooManyElements"
	TariffChangeStatusConditionNotSupported TariffChangeStatusEnumType = "ConditionNotSupported"
	TariffChangeStatusTxNotFound            TariffChangeStatusEnumType = "TxNotFound"
	TariffChangeStatusNoCurrencyChange      TariffChangeStatusEnumType = "NoCurrencyChange"
)

func isValidTariffChangeStatus(fl validator.FieldLevel) bool {
	status := TariffChangeStatusEnumType(fl.Field().String())
	switch status {
	case TariffChangeStatusAccepted, TariffChangeStatusRejected, TariffChangeStatusTooManyElements, TariffChangeStatusConditionNotSupported, TariffChangeStatusTxNotFound, TariffChangeStatusNoCurrencyChange:
		return true
	default:
		return false
	}
}

// The field definition of the ChangeTransactionTariff request payload sent by the CSMS to the Charging Station.
type ChangeTransactionTariffRequest struct {
	Tariff        types.TariffType `json:"tariff" validate:"required"`
	TransactionID string           `json:"transactionId" validate:"required,max=36"`
}

// This field definition of the ChangeTransactionTariff response payload, sent by the Charging Station to the CSMS in response to a ChangeTransactionTariffRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type ChangeTransactionTariffResponse struct {
	Status     TariffChangeStatusEnumType `json:"status" validate:"required,tariffChangeStatus"`
	StatusInfo *types.StatusInfo          `json:"statusInfo,omitempty" validate:"omitempty"`
}

// The driver wants to know how much the running total cost is, updated at a relevant interval, while a transaction is ongoing.
// To fulfill this requirement, the CSMS sends a ChangeTransactionTariffRequest to the Charging Station to update the current total cost, every Y seconds.
// Upon receipt of the ChangeTransactionTariffRequest, the Charging Station responds with a ChangeTransactionTariffResponse, then shows the updated cost to the driver.
type ChangeTransactionTariffFeature struct{}

func (f ChangeTransactionTariffFeature) GetFeatureName() string {
	return ChangeTransactionTariffFeatureName
}

func (f ChangeTransactionTariffFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(ChangeTransactionTariffRequest{})
}

func (f ChangeTransactionTariffFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(ChangeTransactionTariffResponse{})
}

func (r ChangeTransactionTariffRequest) GetFeatureName() string {
	return ChangeTransactionTariffFeatureName
}

func (c ChangeTransactionTariffResponse) GetFeatureName() string {
	return ChangeTransactionTariffFeatureName
}

// Creates a new ChangeTransactionTariffRequest, containing all required fields. There are no optional fields for this message.
func NewChangeTransactionTariffRequest(tariff types.TariffType, transactionID string) *ChangeTransactionTariffRequest {
	return &ChangeTransactionTariffRequest{Tariff: tariff, TransactionID: transactionID}
}

// Creates a new ChangeTransactionTariffResponse, which doesn't contain any required or optional fields.
func NewChangeTransactionTariffResponse(status TariffChangeStatusEnumType) *ChangeTransactionTariffResponse {
	return &ChangeTransactionTariffResponse{Status: status}
}

func init() {
	_ = types.Validate.RegisterValidation("tariffChangeStatus", isValidTariffChangeStatus)
}
