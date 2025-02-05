package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Notify Charging Limit (CS -> CSMS) --------------------

const GetDERControlFeatureName = "GetDERControl"

// The field definition of the GetDERControl request payload sent by the Charging Station to the CSMS.
type GetDERControlRequest struct {
	RequestId   int                       `json:"requestId"`
	IsDefault   bool                      `json:"isDefault,omitempty" validate:"omitempty"`
	ControlType *types.DERControlEnumType `json:"controlType,omitempty" validate:"omitempty,derControlType"`
	ControlId   string                    `json:"controlId,omitempty" validate:"omitempty,max=36"`
}

// This field definition of the GetDERControl response payload, sent by the CSMS to the Charging Station in response to a GetDERControlRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type GetDERControlResponse struct {
	Status     DERControlStatusEnumType `json:"status" validate:"required,derControlStatusType"`
	StatusInfo *types.StatusInfo        `json:"statusInfo,omitempty" validate:"omitempty"`
}

// The CSMS responds with GetDERControlResponse to the Charging Station.
type GetDERControlFeature struct{}

func (f GetDERControlFeature) GetFeatureName() string {
	return GetDERControlFeatureName
}

func (f GetDERControlFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(GetDERControlRequest{})
}

func (f GetDERControlFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(GetDERControlResponse{})
}

func (r GetDERControlRequest) GetFeatureName() string {
	return GetDERControlFeatureName
}

func (c GetDERControlResponse) GetFeatureName() string {
	return GetDERControlFeatureName
}

// Creates a new GetDERControlRequest, containing all required fields. Optional fields may be set afterwards.
func NewGetDERControlRequest(requestId int) *GetDERControlRequest {
	return &GetDERControlRequest{RequestId: requestId}
}

// Creates a new GetDERControlResponse, which doesn't contain any required or optional fields.
func NewGetDERControlResponse() *GetDERControlResponse {
	return &GetDERControlResponse{}
}
