package diagnostics

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"reflect"
)

// -------------------- Notify Settlement (CS -> CSMS) --------------------

const NotifySettlementFeatureName = "NotifySettlement"

// The field definition of the NotifySettlement request payload sent by a Charging Station to the CSMS.
type NotifySettlementRequest struct {
	TransactionId    string                      `json:"transactionId,omitempty" validate:"omitempty,max=36"`
	PPSRef           string                      `json:"pspRef" validate:"required,max=255"`
	Status           types.PaymentStatusEnumType `json:"status" validate:"required,paymentStatusType"`
	StatusInfo       string                      `json:"statusInfo,omitempty" validate:"omitempty,max=500"`
	SettlementAmount float64                     `json:"settlementAmount"`
	SettlementTime   *types.DateTime             `json:"settlementTime" validate:"required"`
	ReceiptId        string                      `json:"receiptId,omitempty" validate:"omitempty,max=50"`
	ReceiptUrl       string                      `json:"receiptUrl,omitempty" validate:"omitempty,max=2000"`
	VatCompany       *types.AddressType          `json:"vatCompany,omitempty" validate:"omitempty"`
	VatNumber        string                      `json:"vatNumber,omitempty" validate:"omitempty,max=20"`
}

// This field definition of the NotifySettlement response payload, sent by the CSMS to the Charging Station in response to a NotifySettlementRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type NotifySettlementResponse struct {
	ReceiptUrl string `json:"receiptUrl,omitempty" validate:"omitempty,max=2000"`
	ReceiptId  string `json:"receiptId,omitempty" validate:"omitempty,max=50"`
}

// The CSMS responds with a NotifySettlementResponse for every received received request.
type NotifySettlementFeature struct{}

func (f NotifySettlementFeature) GetFeatureName() string {
	return NotifySettlementFeatureName
}

func (f NotifySettlementFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(NotifySettlementRequest{})
}

func (f NotifySettlementFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(NotifySettlementResponse{})
}

func (r NotifySettlementRequest) GetFeatureName() string {
	return NotifySettlementFeatureName
}

func (c NotifySettlementResponse) GetFeatureName() string {
	return NotifySettlementFeatureName
}

// Creates a new NotifySettlementRequest, containing required fields. Optional fields may be set afterwards.
func NewNotifySettlementRequest(pspRef string, status types.PaymentStatusEnumType, settlementAmount float64, settlementTime *types.DateTime) *NotifySettlementRequest {
	return &NotifySettlementRequest{PPSRef: pspRef, Status: status, SettlementAmount: settlementAmount, SettlementTime: settlementTime}
}

// Creates a new NotifySettlementResponse, which doesn't contain any required or optional fields.
func NewNotifySettlementResponse() *NotifySettlementResponse {
	return &NotifySettlementResponse{}
}
