package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Set Charging Profile (CSMS -> CS) --------------------

const ClearDERControlFeatureName = "ClearDERControl"

// The field definition of the ClearDERControl request payload sent by the CSMS to the Charging Station.
type ClearDERControlRequest struct {
	IsDefault   bool                     `json:"isDefault"`
	ControlType types.DERControlEnumType `json:"controlType,omitempty" validate:"omitempty"`
	ControlId   string                   `json:"controlId,omitempty" validate:"omitempty,max=36"`
}

// This field definition of the ClearDERControl response payload, sent by the Charging Station to the CSMS in response to a ClearDERControlRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type ClearDERControlResponse struct {
	Status     DERControlStatusEnumType `json:"status" validate:"required,derControlStatusType"`
	StatusInfo *types.StatusInfo        `json:"statusInfo,omitempty" validate:"omitempty"`
}

type ClearDERControlFeature struct{}

func (f ClearDERControlFeature) GetFeatureName() string {
	return ClearDERControlFeatureName
}

func (f ClearDERControlFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(ClearDERControlRequest{})
}

func (f ClearDERControlFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(ClearDERControlResponse{})
}

func (r ClearDERControlRequest) GetFeatureName() string {
	return ClearDERControlFeatureName
}

func (c ClearDERControlResponse) GetFeatureName() string {
	return ClearDERControlFeatureName
}

// Creates a new ClearDERControlRequest, containing all required fields. There are no optional fields for this message.
func NewClearDERControlRequest(isDefault bool) *ClearDERControlRequest {
	return &ClearDERControlRequest{
		IsDefault: isDefault,
	}
}

// Creates a new ClearDERControlResponse, containing all required fields. Optional fields may be set afterwards.
func NewClearDERControlResponse(status DERControlStatusEnumType) *ClearDERControlResponse {
	return &ClearDERControlResponse{Status: status}
}
