package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Notify Charging Limit (CS -> CSMS) --------------------

const NotifyDERStartStopFeatureName = "NotifyDERStartStop"

// The field definition of the NotifyDERStartStop request payload sent by the Charging Station to the CSMS.
type NotifyDERStartStopRequest struct {
	ControlId     string          `json:"controlId" validate:"required,max=36"`
	Started       bool            `json:"started"`
	Timestamp     *types.DateTime `json:"timestamp" validate:"required"`
	SupersededIds []string        `json:"supersededIds,omitempty" validate:"omitempty,min=1,max=24,dive,max=36"`
}

// This field definition of the NotifyDERStartStop response payload, sent by the CSMS to the Charging Station in response to a NotifyDERStartStopRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type NotifyDERStartStopResponse struct {
}

// The CSMS responds with NotifyDERStartStopResponse to the Charging Station.
type NotifyDERStartStopFeature struct{}

func (f NotifyDERStartStopFeature) GetFeatureName() string {
	return NotifyDERStartStopFeatureName
}

func (f NotifyDERStartStopFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(NotifyDERStartStopRequest{})
}

func (f NotifyDERStartStopFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(NotifyDERStartStopResponse{})
}

func (r NotifyDERStartStopRequest) GetFeatureName() string {
	return NotifyDERStartStopFeatureName
}

func (c NotifyDERStartStopResponse) GetFeatureName() string {
	return NotifyDERStartStopFeatureName
}

// Creates a new NotifyDERStartStopRequest, containing all required fields. Optional fields may be set afterwards.
func NewNotifyDERStartStopRequest(controlId string, started bool, timestamp *types.DateTime) *NotifyDERStartStopRequest {
	return &NotifyDERStartStopRequest{ControlId: controlId, Started: started, Timestamp: timestamp}
}

// Creates a new NotifyDERStartStopResponse, which doesn't contain any required or optional fields.
func NewNotifyDERStartStopResponse() *NotifyDERStartStopResponse {
	return &NotifyDERStartStopResponse{}
}
