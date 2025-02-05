package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Notify Charging Limit (CS -> CSMS) --------------------

const NotifyDERAlarmFeatureName = "NotifyDERAlarm"

// The field definition of the NotifyDERAlarm request payload sent by the Charging Station to the CSMS.
type NotifyDERAlarmRequest struct {
	ControlType    types.DERControlEnumType     `json:"controlType" validate:"required,derControlType"`
	GridEventFault types.GridEventFaultEnumType `json:"gridEventFault,omitempty" validate:"omitempty,gridEventFaultType"`
	AlarmEnded     bool                         `json:"alarmEnded,omitempty" validate:"omitempty"`
	Timestamp      *types.DateTime              `json:"timestamp" validate:"required"`
	ExtraInfo      string                       `json:"extraInfo,omitempty" validate:"omitempty,max=200"`
}

// This field definition of the NotifyDERAlarm response payload, sent by the CSMS to the Charging Station in response to a NotifyDERAlarmRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type NotifyDERAlarmResponse struct {
}

// The CSMS responds with NotifyDERAlarmResponse to the Charging Station.
type NotifyDERAlarmFeature struct{}

func (f NotifyDERAlarmFeature) GetFeatureName() string {
	return NotifyDERAlarmFeatureName
}

func (f NotifyDERAlarmFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(NotifyDERAlarmRequest{})
}

func (f NotifyDERAlarmFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(NotifyDERAlarmResponse{})
}

func (r NotifyDERAlarmRequest) GetFeatureName() string {
	return NotifyDERAlarmFeatureName
}

func (c NotifyDERAlarmResponse) GetFeatureName() string {
	return NotifyDERAlarmFeatureName
}

// Creates a new NotifyDERAlarmRequest, containing all required fields. Optional fields may be set afterwards.
func NewNotifyDERAlarmRequest(controlType types.DERControlEnumType, timestamp *types.DateTime) *NotifyDERAlarmRequest {
	return &NotifyDERAlarmRequest{ControlType: controlType, Timestamp: timestamp}
}

// Creates a new NotifyDERAlarmResponse, which doesn't contain any required or optional fields.
func NewNotifyDERAlarmResponse() *NotifyDERAlarmResponse {
	return &NotifyDERAlarmResponse{}
}
