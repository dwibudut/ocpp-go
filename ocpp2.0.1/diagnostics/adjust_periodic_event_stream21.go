package diagnostics

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"reflect"
)

// -------------------- Notify Web Payment Started (CSMS -> CS) --------------------

const AdjustPeriodicEventStreamFeatureName = "AdjustPeriodicEventStream"

// The field definition of the AdjustPeriodicEventStream request payload sent by a Charging Station to the CSMS.
type AdjustPeriodicEventStreamRequest struct {
	Id     int                           `json:"id" validate:"gte=0"`
	Params PeriodicEventStreamParamsType `json:"params" validate:"required"`
}

// This field definition of the AdjustPeriodicEventStream response payload, sent by the CSMS to the Charging Station in response to a AdjustPeriodicEventStreamRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type AdjustPeriodicEventStreamResponse struct {
	Status     types.GenericStatus `json:"status" validate:"required,genericStatus"`
	StatusInfo *types.StatusInfo   `json:"statusInfo,omitempty" validate:"omitempty"`
}

// The CSMS responds with a AdjustPeriodicEventStreamResponse for every received received request.
type AdjustPeriodicEventStreamFeature struct{}

func (f AdjustPeriodicEventStreamFeature) GetFeatureName() string {
	return AdjustPeriodicEventStreamFeatureName
}

func (f AdjustPeriodicEventStreamFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(AdjustPeriodicEventStreamRequest{})
}

func (f AdjustPeriodicEventStreamFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(AdjustPeriodicEventStreamResponse{})
}

func (r AdjustPeriodicEventStreamRequest) GetFeatureName() string {
	return AdjustPeriodicEventStreamFeatureName
}

func (c AdjustPeriodicEventStreamResponse) GetFeatureName() string {
	return AdjustPeriodicEventStreamFeatureName
}

// Creates a new AdjustPeriodicEventStreamRequest, containing all required fields. Optional fields may be set afterwards.
func NewAdjustPeriodicEventStreamRequest(id int, params PeriodicEventStreamParamsType) *AdjustPeriodicEventStreamRequest {
	return &AdjustPeriodicEventStreamRequest{Id: id, Params: params}
}

// Creates a new AdjustPeriodicEventStreamResponse, which doesn't contain any required or optional fields.
func NewAdjustPeriodicEventStreamResponse(status types.GenericStatus) *AdjustPeriodicEventStreamResponse {
	return &AdjustPeriodicEventStreamResponse{Status: status}
}
