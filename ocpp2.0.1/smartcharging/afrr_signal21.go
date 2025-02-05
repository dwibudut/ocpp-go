package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Set Charging Profile (CSMS -> CS) --------------------

const AFRRSignalFeatureName = "AFRRSignal"

// The field definition of the AFRRSignal request payload sent by the CSMS to the Charging Station.
type AFRRSignalRequest struct {
	Timestamp *types.DateTime `json:"timestamp" validate:"required"`
	Signal    int             `json:"signal"`
}

// This field definition of the AFRRSignal response payload, sent by the Charging Station to the CSMS in response to a AFRRSignalRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type AFRRSignalResponse struct {
	Status     types.GenericStatus `json:"status" validate:"required,genericStatus"`
	StatusInfo *types.StatusInfo   `json:"statusInfo,omitempty" validate:"omitempty"`
}

// CSMS sends AFRRSignalRequest to the Charging Station. The charging schedule limits may be imposed by any
// external system. The Charging Station responds to this request with a AFRRSignalResponse.
type AFRRSignalFeature struct{}

func (f AFRRSignalFeature) GetFeatureName() string {
	return AFRRSignalFeatureName
}

func (f AFRRSignalFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(AFRRSignalRequest{})
}

func (f AFRRSignalFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(AFRRSignalResponse{})
}

func (r AFRRSignalRequest) GetFeatureName() string {
	return AFRRSignalFeatureName
}

func (c AFRRSignalResponse) GetFeatureName() string {
	return AFRRSignalFeatureName
}

// Creates a new AFRRSignalRequest, containing all required fields. There are no optional fields for this message.
func NewAFRRSignalRequest(timestamp *types.DateTime, signal int) *AFRRSignalRequest {
	return &AFRRSignalRequest{
		Timestamp: timestamp,
		Signal:    signal,
	}
}

// Creates a new AFRRSignalResponse, containing all required fields. Optional fields may be set afterwards.
func NewAFRRSignalResponse(status types.GenericStatus) *AFRRSignalResponse {
	return &AFRRSignalResponse{Status: status}
}
