package transactions

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Transaction Event (CS -> CSMS) --------------------

const BatterySwapFeatureName = "BatterySwap"

// The type of a transaction event.
type BatterySwapEventEnumType string

const (
	BatterySwapEventBatteryIn         BatterySwapEventEnumType = "BatteryIn"
	BatterySwapEventBatteryOut        BatterySwapEventEnumType = "BatteryOut"
	BatterySwapEventBatteryOutTimeout BatterySwapEventEnumType = "BatteryOutTimeout"
)

func isValidBatterySwapEventType(fl validator.FieldLevel) bool {
	status := BatterySwapEventEnumType(fl.Field().String())
	switch status {
	case BatterySwapEventBatteryIn, BatterySwapEventBatteryOut, BatterySwapEventBatteryOutTimeout:
		return true
	default:
		return false
	}
}

// BatteryDataType
type BatteryDataType struct {
	EvseID         int             `json:"evseId" validate:"gte=0"`
	SerialNumber   string          `json:"serialNumber" validate:"required,max=50"`
	SoC            float64         `json:"soC" validate:"gte=0,lte=100"`
	SoH            float64         `json:"soH" validate:"gte=0,lte=100"`
	ProductionDate *types.DateTime `json:"productionDate,omitempty" validate:"omitempty"`
	VendorInfo     string          `json:"vendorInfo,omitempty" validate:"omitempty,max=500"`
}

// The field definition of the BatterySwap request payload sent by the Charging Station to the CSMS.
type BatterySwapRequest struct {
	BatteryData []BatteryDataType        `json:"batteryData" validate:"required,min=1,dive"`
	EventType   BatterySwapEventEnumType `json:"eventType" validate:"required,batterySwapEventType"`
	IDToken     *types.IdToken           `json:"idToken" validate:"required"`
	RequestId   int                      `json:"requestId"`
}

// This field definition of the BatterySwapResponse payload, sent by the CSMS to the Charging Station in response to a BatterySwapRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type BatterySwapResponse struct {
}

// Gives the CSMS information that will later be used to bill a transaction.
// For this purpose, status changes and additional transaction-related information is sent, such as
// retrying and sequence number messages.
//
// A Charging Station notifies the CSMS using a BatterySwapRequest. The CSMS then responds with a
// BatterySwapResponse.
type BatterySwapFeature struct{}

func (f BatterySwapFeature) GetFeatureName() string {
	return BatterySwapFeatureName
}

func (f BatterySwapFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(BatterySwapRequest{})
}

func (f BatterySwapFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(BatterySwapResponse{})
}

func (r BatterySwapRequest) GetFeatureName() string {
	return BatterySwapFeatureName
}

func (c BatterySwapResponse) GetFeatureName() string {
	return BatterySwapFeatureName
}

// Creates a new BatterySwapRequest, containing all required fields. Optional fields may be set afterwards.
func NewBatterySwapRequest(batteryData []BatteryDataType, eventType BatterySwapEventEnumType, idToken *types.IdToken, requestId int) *BatterySwapRequest {
	return &BatterySwapRequest{BatteryData: batteryData, EventType: eventType, IDToken: idToken, RequestId: requestId}
}

// Creates a new BatterySwapResponse, containing all required fields. Optional fields may be set afterwards.
func NewBatterySwapResponse() *BatterySwapResponse {
	return &BatterySwapResponse{}
}

func init() {
	_ = types.Validate.RegisterValidation("batterySwapEventType", isValidBatterySwapEventType)
}
