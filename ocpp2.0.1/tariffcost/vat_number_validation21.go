package tariffcost

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"reflect"
)

// -------------------- Cost Updated (CSMS -> CS) --------------------

const VatNumberValidationFeatureName = "VatNumberValidation"

// The field definition of the VatNumberValidation request payload sent by the CSMS to the Charging Station.
type VatNumberValidationRequest struct {
	VatNumber string `json:"vatNumber" validate:"required,max=20"`
	EvseId    int    `json:"evseId,omitempty" validate:"omitempty,gte=0"`
}

// This field definition of the VatNumberValidation response payload, sent by the Charging Station to the CSMS in response to a VatNumberValidationRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type VatNumberValidationResponse struct {
	Company    *types.AddressType  `json:"company,omitempty" validate:"omitempty"`
	StatusInfo *types.StatusInfo   `json:"statusInfo,omitempty" validate:"omitempty"`
	VatNumber  string              `json:"vatNumber" validate:"required,max=20"`
	EvseId     int                 `json:"evseId,omitempty" validate:"omitempty,gte=0"`
	Status     types.GenericStatus `json:"status" validate:"required,genericStatus"`
}

// The driver wants to know how much the running total cost is, updated at a relevant interval, while a transaction is ongoing.
// To fulfill this requirement, the CSMS sends a VatNumberValidationRequest to the Charging Station to update the current total cost, every Y seconds.
// Upon receipt of the VatNumberValidationRequest, the Charging Station responds with a VatNumberValidationResponse, then shows the updated cost to the driver.
type VatNumberValidationFeature struct{}

func (f VatNumberValidationFeature) GetFeatureName() string {
	return VatNumberValidationFeatureName
}

func (f VatNumberValidationFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(VatNumberValidationRequest{})
}

func (f VatNumberValidationFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(VatNumberValidationResponse{})
}

func (r VatNumberValidationRequest) GetFeatureName() string {
	return VatNumberValidationFeatureName
}

func (c VatNumberValidationResponse) GetFeatureName() string {
	return VatNumberValidationFeatureName
}

// Creates a new VatNumberValidationRequest, containing required fields
func NewVatNumberValidationRequest(vatNumber string) *VatNumberValidationRequest {
	return &VatNumberValidationRequest{VatNumber: vatNumber}
}

// Creates a new VatNumberValidationResponse, containing all required fields. There are no optional fields for this message.
func NewVatNumberValidationResponse(vatNumber string, status types.GenericStatus) *VatNumberValidationResponse {
	return &VatNumberValidationResponse{
		VatNumber: vatNumber,
		Status:    status,
	}
}
