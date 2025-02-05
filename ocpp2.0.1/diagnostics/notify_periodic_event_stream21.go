package diagnostics

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Notify Periodic Event Stream (CS -> CSMS) --------------------

const NotifyPeriodicEventStreamFeatureName = "NotifyPeriodicEventStream"

// A StreamDataElementType element contains only the Component, Variable and VariableMonitoring data that caused an event.
type StreamDataElementType struct {
	T float64 `json:"t" validate:"gte=0"`
	V string  `json:"v" validate:"min=0,max=2500"`
}

// The field definition of the NotifyPeriodicEventStream request payload sent by a Charging Station to the CSMS.
type NotifyPeriodicEventStreamRequest struct {
	ID       int                     `json:"id" validate:"gte=0"`
	Pending  int                     `json:"pending" validate:"gte=0"`
	BaseTime *types.DateTime         `json:"basetime" validate:"required"`
	Data     []StreamDataElementType `json:"data" validate:"required,min=1,dive"`
}

type NotifyPeriodicEventStreamFeature struct{}

func (f NotifyPeriodicEventStreamFeature) GetFeatureName() string {
	return NotifyPeriodicEventStreamFeatureName
}

func (f NotifyPeriodicEventStreamFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(NotifyPeriodicEventStreamRequest{})
}

func (f NotifyPeriodicEventStreamFeature) GetResponseType() reflect.Type {
	return nil
}

func (r NotifyPeriodicEventStreamRequest) GetFeatureName() string {
	return NotifyPeriodicEventStreamFeatureName
}

// Creates a new NotifyPeriodicEventStreamRequest, containing all required fields. Optional fields may be set afterwards.
func NewNotifyPeriodicEventStreamRequest(id, pending int, basetime *types.DateTime, data []StreamDataElementType) *NotifyPeriodicEventStreamRequest {
	return &NotifyPeriodicEventStreamRequest{
		ID:       id,
		Pending:  pending,
		BaseTime: basetime,
		Data:     data,
	}
}
