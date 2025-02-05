package diagnostics

import (
	"reflect"
)

// -------------------- Notify Priority Charging (CS -> CSMS) --------------------

const ClosePeriodicEventStreamFeatureName = "ClosePeriodicEventStream"

// The field definition of the ClosePeriodicEventStream request payload sent by a Charging Station to the CSMS.
type ClosePeriodicEventStreamRequest struct {
	Id int `json:"id" validate:"gte=0"`
}

// This field definition of the ClosePeriodicEventStream response payload, sent by the CSMS to the Charging Station in response to a ClosePeriodicEventStreamRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type ClosePeriodicEventStreamResponse struct {
}

// The CSMS responds with a ClosePeriodicEventStreamResponse for every received received request.
type ClosePeriodicEventStreamFeature struct{}

func (f ClosePeriodicEventStreamFeature) GetFeatureName() string {
	return ClosePeriodicEventStreamFeatureName
}

func (f ClosePeriodicEventStreamFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(ClosePeriodicEventStreamRequest{})
}

func (f ClosePeriodicEventStreamFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(ClosePeriodicEventStreamResponse{})
}

func (r ClosePeriodicEventStreamRequest) GetFeatureName() string {
	return ClosePeriodicEventStreamFeatureName
}

func (c ClosePeriodicEventStreamResponse) GetFeatureName() string {
	return ClosePeriodicEventStreamFeatureName
}

// Creates a new ClosePeriodicEventStreamRequest, containing all required fields. Optional fields may be set afterwards.
func NewClosePeriodicEventStreamRequest(id int) *ClosePeriodicEventStreamRequest {
	return &ClosePeriodicEventStreamRequest{Id: id}
}

// Creates a new ClosePeriodicEventStreamResponse, which doesn't contain any required or optional fields.
func NewClosePeriodicEventStreamResponse() *ClosePeriodicEventStreamResponse {
	return &ClosePeriodicEventStreamResponse{}
}
