package smartcharging

import (
	"reflect"
)

// -------------------- Report Charging Profiles (CS -> CSMS) --------------------

const ReportDERControlFeatureName = "ReportDERControl"

// The field definition of the ReportDERControl request payload sent by the Charging Station to the CSMS.
type ReportDERControlRequest struct {
	Curve             []DERCurveGetType          `json:"curve,omitempty" validate:"omitempty,min=1,max=24,dive"`
	EnterService      []EnterServiceGetType      `json:"enterService,omitempty" validate:"omitempty,min=1,max=24,dive"`
	FixedPFAbsorb     []FixedPFGetType           `json:"fixedPFAbsorb,omitempty" validate:"omitempty,min=1,max=24,dive"`
	FixedPFInject     []FixedPFGetType           `json:"fixedPFInject,omitempty" validate:"omitempty,min=1,max=24,dive"`
	FixedVar          []FixedVarGetType          `json:"fixedVar,omitempty" validate:"omitempty,min=1,max=24,dive"`
	FreqDroop         []FreqDroopGetType         `json:"freqDroop,omitempty" validate:"omitempty,min=1,max=24,dive"`
	Gradient          []GradientGetType          `json:"gradient,omitempty" validate:"omitempty,min=1,max=24,dive"`
	LimitMaxDischarge []LimitMaxDischargeGetType `json:"limitMaxDischarge,omitempty" validate:"omitempty,min=1,max=24,dive"`
	RequestId         int                        `json:"requestId"`
	Tbc               bool                       `json:"tbc,omitempty" validate:"omitempty"`
}

// This field definition of the ReportDERControl response payload, sent by the CSMS to the Charging Station in
// response to a ReportDERControlRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type ReportDERControlResponse struct {
}

type ReportDERControlFeature struct{}

func (f ReportDERControlFeature) GetFeatureName() string {
	return ReportDERControlFeatureName
}

func (f ReportDERControlFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(ReportDERControlRequest{})
}

func (f ReportDERControlFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(ReportDERControlResponse{})
}

func (r ReportDERControlRequest) GetFeatureName() string {
	return ReportDERControlFeatureName
}

func (c ReportDERControlResponse) GetFeatureName() string {
	return ReportDERControlFeatureName
}

// Creates a new ReportDERControlRequest, containing all required fields. Optional fields may be set afterwards.
func NewReportDERControlRequest(requestId int) *ReportDERControlRequest {
	return &ReportDERControlRequest{RequestId: requestId}
}

// Creates a new ReportDERControlResponse, which doesn't contain any required or optional fields.
func NewReportDERControlResponse() *ReportDERControlResponse {
	return &ReportDERControlResponse{}
}
