package tariffcost

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// -------------------- Cost Updated (CSMS -> CS) --------------------

const GetTariffsFeatureName = "GetTariffs"

type TariffGetStatusEnumType string

const (
	TariffGetStatusAccepted TariffGetStatusEnumType = "Accepted"
	TariffGetStatusRejected TariffGetStatusEnumType = "Rejected"
	TariffGetStatusNoTariff TariffGetStatusEnumType = "NoTariff"
)

func isValidTariffGetStatus(fl validator.FieldLevel) bool {
	status := TariffGetStatusEnumType(fl.Field().String())
	switch status {
	case TariffGetStatusAccepted, TariffGetStatusRejected, TariffGetStatusNoTariff:
		return true
	default:
		return false
	}
}

// TariffKindEnumType
type TariffKindEnumType string

const (
	TariffKindDefaultTariff TariffKindEnumType = "DefaultTariff"
	TariffKindDriverTariff  TariffKindEnumType = "DriverTariff"
)

func isValidTariffKindType(fl validator.FieldLevel) bool {
	status := TariffKindEnumType(fl.Field().String())
	switch status {
	case TariffKindDefaultTariff, TariffKindDriverTariff:
		return true
	default:
		return false
	}
}

// TariffAssignmentType
type TariffAssignmentType struct {
	TariffId   string             `json:"tariffId" validate:"required,max=60"`
	TariffKind TariffKindEnumType `json:"tariffKind" validate:"required,tariffKindType"`
	ValidFrom  *types.DateTime    `json:"validFrom,omitempty" validate:"omitempty"`
	EvseIds    []int              `json:"evseIds,omitempty" validate:"omitempty,min=1,dive,gte=0"`
	IdTokens   []string           `json:"idTokens,omitempty" validate:"omitempty,min=1,dive,max=255"`
}

// The field definition of the GetTariffs request payload sent by the CSMS to the Charging Station.
type GetTariffsRequest struct {
	EvseId int `json:"evseId" validate:"gte=0"`
}

// This field definition of the GetTariffs response payload, sent by the Charging Station to the CSMS in response to a GetTariffsRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type GetTariffsResponse struct {
	Status            TariffGetStatusEnumType `json:"status" validate:"required,tariffGetStatus"`
	StatusInfo        *types.StatusInfo       `json:"statusInfo,omitempty" validate:"omitempty"`
	TariffAssignments []TariffAssignmentType  `json:"tariffAssignments,omitempty" validate:"omitempty,min=1,dive"`
}

type GetTariffsFeature struct{}

func (f GetTariffsFeature) GetFeatureName() string {
	return GetTariffsFeatureName
}

func (f GetTariffsFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(GetTariffsRequest{})
}

func (f GetTariffsFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(GetTariffsResponse{})
}

func (r GetTariffsRequest) GetFeatureName() string {
	return GetTariffsFeatureName
}

func (c GetTariffsResponse) GetFeatureName() string {
	return GetTariffsFeatureName
}

// Creates a new GetTariffsRequest, containing all required fields. There are no optional fields for this message.
func NewGetTariffsRequest(evseId int) *GetTariffsRequest {
	return &GetTariffsRequest{EvseId: evseId}
}

// Creates a new GetTariffsResponse, which doesn't contain any required or optional fields.
func NewGetTariffsResponse(status TariffGetStatusEnumType) *GetTariffsResponse {
	return &GetTariffsResponse{Status: status}
}

func init() {
	_ = types.Validate.RegisterValidation("tariffGetStatus", isValidTariffGetStatus)
	_ = types.Validate.RegisterValidation("tariffKindType", isValidTariffKindType)
}
