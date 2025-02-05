package iso15118

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Get Certificate Status (CS -> CSMS) --------------------

const GetCertificateChainStatusFeatureName = "GetCertificateChainStatus"

// CertificateStatusRequestInfoType
type CertificateStatusRequestInfoType struct {
	CertificateHashData types.CertificateHashData             `json:"certificateHashData" validate:"required"`
	Source              types.CertificateStatusSourceEnumType `json:"source" validate:"required,certificateStatusSource"`
	Urls                []string                              `json:"urls" validate:"required,min=1,max=5,dive,max=2000"`
}

// The field definition of the GetCertificateChainStatus request payload sent by the Charging Station to the CSMS.
type GetCertificateChainStatusRequest struct {
	CertificateStatusRequests []CertificateStatusRequestInfoType `json:"certificateStatusRequests" validate:"required,min=1,max=4,dive"`
}

// This field definition of the GetCertificateChainStatus response payload, sent by the CSMS to the Charging Station in response to a GetCertificateChainStatusRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type GetCertificateChainStatusResponse struct {
	CertificateStatus []types.CertificateStatusType `json:"certificateStatus" validate:"required,min=1,max=4,dive"`
}

type GetCertificateChainStatusFeature struct{}

func (f GetCertificateChainStatusFeature) GetFeatureName() string {
	return GetCertificateChainStatusFeatureName
}

func (f GetCertificateChainStatusFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(GetCertificateChainStatusRequest{})
}

func (f GetCertificateChainStatusFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(GetCertificateChainStatusResponse{})
}

func (r GetCertificateChainStatusRequest) GetFeatureName() string {
	return GetCertificateChainStatusFeatureName
}

func (c GetCertificateChainStatusResponse) GetFeatureName() string {
	return GetCertificateChainStatusFeatureName
}

func NewCertificateStatusRequestInfo(certificateHashData types.CertificateHashData, source types.CertificateStatusSourceEnumType, urls []string) *CertificateStatusRequestInfoType {
	return &CertificateStatusRequestInfoType{CertificateHashData: certificateHashData, Source: source, Urls: urls}
}

// Creates a new GetCertificateChainStatusRequest, containing all required fields. There are no optional fields for this message.
func NewGetCertificateChainStatusRequest(certificateStatusRequests []CertificateStatusRequestInfoType) *GetCertificateChainStatusRequest {
	return &GetCertificateChainStatusRequest{CertificateStatusRequests: certificateStatusRequests}
}

// Creates a new GetCertificateChainStatusResponse, containing all required fields. Optional fields may be set afterwards.
func NewGetCertificateChainStatusResponse(certificateStatus []types.CertificateStatusType) *GetCertificateChainStatusResponse {
	return &GetCertificateChainStatusResponse{CertificateStatus: certificateStatus}
}
