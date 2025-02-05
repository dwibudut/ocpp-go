package diagnostics

import (
	"reflect"
)

// -------------------- Notify Priority Charging (CS -> CSMS) --------------------

const NotifyPriorityChargingFeatureName = "NotifyPriorityCharging"

// The field definition of the NotifyPriorityCharging request payload sent by a Charging Station to the CSMS.
type NotifyPriorityChargingRequest struct {
	TransactionId string `json:"transactionId" validate:"required,max=36"`
	Activated     bool   `json:"activated"`
}

// This field definition of the NotifyPriorityCharging response payload, sent by the CSMS to the Charging Station in response to a NotifyPriorityChargingRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type NotifyPriorityChargingResponse struct {
}

// The CSMS responds with a NotifyPriorityChargingResponse for every received received request.
type NotifyPriorityChargingFeature struct{}

func (f NotifyPriorityChargingFeature) GetFeatureName() string {
	return NotifyPriorityChargingFeatureName
}

func (f NotifyPriorityChargingFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(NotifyPriorityChargingRequest{})
}

func (f NotifyPriorityChargingFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(NotifyPriorityChargingResponse{})
}

func (r NotifyPriorityChargingRequest) GetFeatureName() string {
	return NotifyPriorityChargingFeatureName
}

func (c NotifyPriorityChargingResponse) GetFeatureName() string {
	return NotifyPriorityChargingFeatureName
}

// Creates a new NotifyPriorityChargingRequest, containing all required fields. Optional fields may be set afterwards.
func NewNotifyPriorityChargingRequest(transactionId string, activated bool) *NotifyPriorityChargingRequest {
	return &NotifyPriorityChargingRequest{TransactionId: transactionId, Activated: activated}
}

// Creates a new NotifyPriorityChargingResponse, which doesn't contain any required or optional fields.
func NewNotifyPriorityChargingResponse() *NotifyPriorityChargingResponse {
	return &NotifyPriorityChargingResponse{}
}
