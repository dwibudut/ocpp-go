package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Set Charging Profile (CSMS -> CS) --------------------

const UsePriorityChargingFeatureName = "UsePriorityCharging"

// The field definition of the UsePriorityCharging request payload sent by the CSMS to the Charging Station.
type UsePriorityChargingRequest struct {
	TransactionId string `json:"transactionId" validate:"required,max=36"`
	Activate      bool   `json:"activate"`
}

// This field definition of the UsePriorityCharging response payload, sent by the Charging Station to the CSMS in response to a UsePriorityChargingRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type UsePriorityChargingResponse struct {
	Status     ChargingProfileStatus `json:"status" validate:"required,chargingProfileStatus201"`
	StatusInfo *types.StatusInfo     `json:"statusInfo,omitempty" validate:"omitempty"`
}

type UsePriorityChargingFeature struct{}

func (f UsePriorityChargingFeature) GetFeatureName() string {
	return UsePriorityChargingFeatureName
}

func (f UsePriorityChargingFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(UsePriorityChargingRequest{})
}

func (f UsePriorityChargingFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(UsePriorityChargingResponse{})
}

func (r UsePriorityChargingRequest) GetFeatureName() string {
	return UsePriorityChargingFeatureName
}

func (c UsePriorityChargingResponse) GetFeatureName() string {
	return UsePriorityChargingFeatureName
}

// Creates a new UsePriorityChargingRequest, containing all required fields. There are no optional fields for this message.
func NewUsePriorityChargingRequest(transactionId string, activate bool) *UsePriorityChargingRequest {
	return &UsePriorityChargingRequest{
		TransactionId: transactionId,
		Activate:      activate,
	}
}

// Creates a new UsePriorityChargingResponse, containing all required fields. Optional fields may be set afterwards.
func NewUsePriorityChargingResponse(status ChargingProfileStatus) *UsePriorityChargingResponse {
	return &UsePriorityChargingResponse{Status: status}
}
