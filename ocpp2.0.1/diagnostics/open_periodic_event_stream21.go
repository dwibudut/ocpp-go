package diagnostics

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"reflect"
)

// -------------------- Open Periodic Event Stream (CS -> CSMS) --------------------

const OpenPeriodicEventStreamFeatureName = "OpenPeriodicEventStream"

// PeriodicEventStreamParamsType
type PeriodicEventStreamParamsType struct {
	Interval int `json:"interval,omitempty" validate:"omitempty,gte=0"`
	Values   int `json:"values,omitempty" validate:"omitempty,gte=0"`
}

// ConstantStreamDataType
type ConstantStreamDataType struct {
	Id                   int                           `json:"id" validate:"gte=0"`
	Params               PeriodicEventStreamParamsType `json:"params" validate:"required"`
	VariableMonitoringId int                           `json:"variableMonitoringId" validate:"gte=0"`
}

// The field definition of the OpenPeriodicEventStream request payload sent by a Charging Station to the CSMS.
type OpenPeriodicEventStreamRequest struct {
	ConstantStreamData ConstantStreamDataType `json:"constantStreamData" validate:"required"`
}

// This field definition of the OpenPeriodicEventStream response payload, sent by the CSMS to the Charging Station in response to a OpenPeriodicEventStreamRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type OpenPeriodicEventStreamResponse struct {
	Status     types.GenericStatus `json:"status" validate:"required,genericStatus"`
	StatusInfo *types.StatusInfo   `json:"statusInfo,omitempty" validate:"omitempty"`
}

// The CSMS responds with a OpenPeriodicEventStreamResponse for every received received request.
type OpenPeriodicEventStreamFeature struct{}

func (f OpenPeriodicEventStreamFeature) GetFeatureName() string {
	return OpenPeriodicEventStreamFeatureName
}

func (f OpenPeriodicEventStreamFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(OpenPeriodicEventStreamRequest{})
}

func (f OpenPeriodicEventStreamFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(OpenPeriodicEventStreamResponse{})
}

func (r OpenPeriodicEventStreamRequest) GetFeatureName() string {
	return OpenPeriodicEventStreamFeatureName
}

func (c OpenPeriodicEventStreamResponse) GetFeatureName() string {
	return OpenPeriodicEventStreamFeatureName
}

// Creates a new OpenPeriodicEventStreamRequest, containing all required fields. Optional fields may be set afterwards.
func NewOpenPeriodicEventStreamRequest(constantStreamData ConstantStreamDataType) *OpenPeriodicEventStreamRequest {
	return &OpenPeriodicEventStreamRequest{ConstantStreamData: constantStreamData}
}

// Creates a new OpenPeriodicEventStreamResponse, which doesn't contain any required or optional fields.
func NewOpenPeriodicEventStreamResponse(status types.GenericStatus) *OpenPeriodicEventStreamResponse {
	return &OpenPeriodicEventStreamResponse{Status: status}
}
