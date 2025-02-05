package remotecontrol

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Request Battery Swap (CSMS -> CS) --------------------

const RequestBatterySwapFeatureName = "RequestBatterySwap"

// The field definition of the RequestBatterySwap request payload sent by the CSMS to the Charging Station.
type RequestBatterySwapRequest struct {
	IDToken   types.IdToken `json:"idToken"`
	RequestId int           `json:"requestId"`
}

// This field definition of the RequestBatterySwap response payload, sent by the Charging Station to the CSMS in response to a RequestBatterySwapRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type RequestBatterySwapResponse struct {
	Status     types.GenericStatus `json:"status" validate:"required,genericStatus"`
	StatusInfo *types.StatusInfo   `json:"statusInfo,omitempty"`
}

// The CSMS may remotely battery swap for a user.
//
// The CSMS sends a RequestBatterySwapRequest to the Charging Station.
// The Charging Stations will reply with a RequestBatterySwapResponse.
type RequestBatterySwapFeature struct{}

func (f RequestBatterySwapFeature) GetFeatureName() string {
	return RequestBatterySwapFeatureName
}

func (f RequestBatterySwapFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(RequestBatterySwapRequest{})
}

func (f RequestBatterySwapFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(RequestBatterySwapResponse{})
}

func (r RequestBatterySwapRequest) GetFeatureName() string {
	return RequestBatterySwapFeatureName
}

func (c RequestBatterySwapResponse) GetFeatureName() string {
	return RequestBatterySwapFeatureName
}

// Creates a new RequestBatterySwapRequest, containing all required fields. Optional fields may be set afterwards.
func NewRequestBatterySwapRequest(requestId int, IdToken types.IdToken) *RequestBatterySwapRequest {
	return &RequestBatterySwapRequest{RequestId: requestId, IDToken: IdToken}
}

// Creates a new RequestBatterySwapResponse, containing all required fields. Optional fields may be set afterwards.
func NewRequestBatterySwapResponse(status types.GenericStatus) *RequestBatterySwapResponse {
	return &RequestBatterySwapResponse{Status: status}
}
