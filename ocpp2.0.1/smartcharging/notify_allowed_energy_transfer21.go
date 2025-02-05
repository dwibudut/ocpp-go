package smartcharging

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Set Charging Profile (CSMS -> CS) --------------------

const NotifyAllowedEnergyTransferFeatureName = "NotifyAllowedEnergyTransfer"

// NotifyAllowedEnergyTransferStatusEnumType
type NotifyAllowedEnergyTransferStatusEnumType string

const (
	NotifyAllowedEnergyTransferStatusAccepted NotifyAllowedEnergyTransferStatusEnumType = "Accepted"
	NotifyAllowedEnergyTransferStatusRejected NotifyAllowedEnergyTransferStatusEnumType = "Rejected"
)

func isNotifyAllowedEnergyTransferStatus(fl validator.FieldLevel) bool {
	status := NotifyAllowedEnergyTransferStatusEnumType(fl.Field().String())
	switch status {
	case NotifyAllowedEnergyTransferStatusAccepted, NotifyAllowedEnergyTransferStatusRejected:
		return true
	default:
		return false
	}
}

// The field definition of the NotifyAllowedEnergyTransfer request payload sent by the CSMS to the Charging Station.
type NotifyAllowedEnergyTransferRequest struct {
	TransactionId         string                             `json:"transactionId" validate:"required,max=36"`
	AllowedEnergyTransfer []types.EnergyTransferModeEnumType `json:"allowedEnergyTransfer" validate:"required,min=1,dive,allowedEnergyTransferStatus"`
}

// This field definition of the NotifyAllowedEnergyTransfer response payload, sent by the Charging Station to the CSMS in response to a NotifyAllowedEnergyTransferRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type NotifyAllowedEnergyTransferResponse struct {
	Status     NotifyAllowedEnergyTransferStatusEnumType `json:"status" validate:"required,chargingProfileStatus201"`
	StatusInfo *types.StatusInfo                         `json:"statusInfo,omitempty" validate:"omitempty"`
}

type NotifyAllowedEnergyTransferFeature struct{}

func (f NotifyAllowedEnergyTransferFeature) GetFeatureName() string {
	return NotifyAllowedEnergyTransferFeatureName
}

func (f NotifyAllowedEnergyTransferFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(NotifyAllowedEnergyTransferRequest{})
}

func (f NotifyAllowedEnergyTransferFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(NotifyAllowedEnergyTransferResponse{})
}

func (r NotifyAllowedEnergyTransferRequest) GetFeatureName() string {
	return NotifyAllowedEnergyTransferFeatureName
}

func (c NotifyAllowedEnergyTransferResponse) GetFeatureName() string {
	return NotifyAllowedEnergyTransferFeatureName
}

// Creates a new NotifyAllowedEnergyTransferRequest, containing all required fields. There are no optional fields for this message.
func NewNotifyAllowedEnergyTransferRequest(transactionId string, allowedEnergyTransfer []types.EnergyTransferModeEnumType) *NotifyAllowedEnergyTransferRequest {
	return &NotifyAllowedEnergyTransferRequest{
		TransactionId:         transactionId,
		AllowedEnergyTransfer: allowedEnergyTransfer,
	}
}

// Creates a new NotifyAllowedEnergyTransferResponse, containing all required fields. Optional fields may be set afterwards.
func NewNotifyAllowedEnergyTransferResponse(status NotifyAllowedEnergyTransferStatusEnumType) *NotifyAllowedEnergyTransferResponse {
	return &NotifyAllowedEnergyTransferResponse{Status: status}
}

func init() {
	_ = types.Validate.RegisterValidation("allowedEnergyTransferStatus", isNotifyAllowedEnergyTransferStatus)
}
