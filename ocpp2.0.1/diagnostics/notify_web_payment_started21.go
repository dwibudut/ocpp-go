package diagnostics

import (
	"reflect"
)

// -------------------- Notify Web Payment Started (CSMS -> CS) --------------------

const NotifyWebPaymentStartedFeatureName = "NotifyWebPaymentStarted"

// The field definition of the NotifyWebPaymentStarted request payload sent by a Charging Station to the CSMS.
type NotifyWebPaymentStartedRequest struct {
	EvseId  int `json:"evseId" validate:"gte=0"`
	Timeout int `json:"timeout"`
}

// This field definition of the NotifyWebPaymentStarted response payload, sent by the CSMS to the Charging Station in response to a NotifyWebPaymentStartedRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type NotifyWebPaymentStartedResponse struct {
}

// The CSMS responds with a NotifyWebPaymentStartedResponse for every received received request.
type NotifyWebPaymentStartedFeature struct{}

func (f NotifyWebPaymentStartedFeature) GetFeatureName() string {
	return NotifyWebPaymentStartedFeatureName
}

func (f NotifyWebPaymentStartedFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(NotifyWebPaymentStartedRequest{})
}

func (f NotifyWebPaymentStartedFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(NotifyWebPaymentStartedResponse{})
}

func (r NotifyWebPaymentStartedRequest) GetFeatureName() string {
	return NotifyWebPaymentStartedFeatureName
}

func (c NotifyWebPaymentStartedResponse) GetFeatureName() string {
	return NotifyWebPaymentStartedFeatureName
}

// Creates a new NotifyWebPaymentStartedRequest, containing all required fields. Optional fields may be set afterwards.
func NewNotifyWebPaymentStartedRequest(evseId int, timeout int) *NotifyWebPaymentStartedRequest {
	return &NotifyWebPaymentStartedRequest{EvseId: evseId, Timeout: timeout}
}

// Creates a new NotifyWebPaymentStartedResponse, which doesn't contain any required or optional fields.
func NewNotifyWebPaymentStartedResponse() *NotifyWebPaymentStartedResponse {
	return &NotifyWebPaymentStartedResponse{}
}
