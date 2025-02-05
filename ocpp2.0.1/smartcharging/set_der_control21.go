package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Set Charging Profile (CSMS -> CS) --------------------

const SetDERControlFeatureName = "SetDERControl"

// The field definition of the SetDERControl request payload sent by the CSMS to the Charging Station.
type SetDERControlRequest struct {
	IsDefault         bool                     `json:"isDefault"`
	ControlId         string                   `json:"controlId" validate:"required,max=36"`
	ControlType       types.DERControlEnumType `json:"controlType" validate:"required"`
	Curve             *DERCurveType            `json:"curve,omitempty" validate:"omitempty"`
	EnterService      *EnterServiceType        `json:"enterService,omitempty" validate:"omitempty"`
	FixedPFAbsorb     *FixedPFType             `json:"fixedPFAbsorb,omitempty" validate:"omitempty"`
	FixedPFInject     *FixedPFType             `json:"fixedPFInject,omitempty" validate:"omitempty"`
	FixedVar          *FixedVarType            `json:"fixedVar,omitempty" validate:"omitempty"`
	FreqDroop         *FreqDroopType           `json:"freqDroop,omitempty" validate:"omitempty"`
	Gradient          *GradientType            `json:"gradient,omitempty" validate:"omitempty"`
	LimitMaxDischarge *LimitMaxDischargeType   `json:"limitMaxDischarge,omitempty" validate:"omitempty"`
}

// This field definition of the SetDERControl response payload, sent by the Charging Station to the CSMS in response to a SetDERControlRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type SetDERControlResponse struct {
	Status        DERControlStatusEnumType `json:"status" validate:"required,derControlStatusType"`
	StatusInfo    *types.StatusInfo        `json:"statusInfo,omitempty" validate:"omitempty"`
	SupersededIds []string                 `json:"supersededIds,omitempty" validate:"omitempty,min=1,max=24,dive,max=36"`
}

type SetDERControlFeature struct{}

func (f SetDERControlFeature) GetFeatureName() string {
	return SetDERControlFeatureName
}

func (f SetDERControlFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(SetDERControlRequest{})
}

func (f SetDERControlFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(SetDERControlResponse{})
}

func (r SetDERControlRequest) GetFeatureName() string {
	return SetDERControlFeatureName
}

func (c SetDERControlResponse) GetFeatureName() string {
	return SetDERControlFeatureName
}

// Creates a new SetDERControlRequest, containing all required fields. There are no optional fields for this message.
func NewSetDERControlRequest(isDefault bool, controlId string, controlType types.DERControlEnumType) *SetDERControlRequest {
	return &SetDERControlRequest{
		IsDefault:   isDefault,
		ControlId:   controlId,
		ControlType: controlType,
	}
}

// Creates a new SetDERControlResponse, containing all required fields. Optional fields may be set afterwards.
func NewSetDERControlResponse(status DERControlStatusEnumType) *SetDERControlResponse {
	return &SetDERControlResponse{Status: status}
}
