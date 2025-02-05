package diagnostics

import (
	"reflect"
)

// -------------------- Get Monitoring Report (CSMS -> CS) --------------------

const GetPeriodicEventStreamFeatureName = "GetPeriodicEventStream"

// The field definition of the GetPeriodicEventStream request payload sent by the CSMS to the Charging Station.
type GetPeriodicEventStreamRequest struct{}

// This field definition of the GetPeriodicEventStream response payload, sent by the Charging Station to the CSMS in response to a GetPeriodicEventStreamRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type GetPeriodicEventStreamResponse struct {
	ConstantStreamData []ConstantStreamDataType `json:"constantStreamData,omitempty" validate:"omitempty,min=1,dive"`
}

type GetPeriodicEventStreamFeature struct{}

func (f GetPeriodicEventStreamFeature) GetFeatureName() string {
	return GetPeriodicEventStreamFeatureName
}

func (f GetPeriodicEventStreamFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(GetPeriodicEventStreamRequest{})
}

func (f GetPeriodicEventStreamFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(GetPeriodicEventStreamResponse{})
}

func (r GetPeriodicEventStreamRequest) GetFeatureName() string {
	return GetPeriodicEventStreamFeatureName
}

func (c GetPeriodicEventStreamResponse) GetFeatureName() string {
	return GetPeriodicEventStreamFeatureName
}

// Creates a new GetPeriodicEventStreamRequest. All fields are optional and may be set afterwards.
func NewGetPeriodicEventStreamRequest() *GetPeriodicEventStreamRequest {
	return &GetPeriodicEventStreamRequest{}
}

// Creates a new GetPeriodicEventStreamResponse, containing all required fields. There are no optional fields for this message.
func NewGetPeriodicEventStreamResponse() *GetPeriodicEventStreamResponse {
	return &GetPeriodicEventStreamResponse{}
}
