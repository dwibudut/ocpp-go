package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Set Charging Profile (CS -> CSMS) --------------------

const PullDynamicScheduleUpdateFeatureName = "PullDynamicScheduleUpdate"

// The field definition of the PullDynamicScheduleUpdate request payload sent by the CSMS to the Charging Station.
type PullDynamicScheduleUpdateRequest struct {
	ChargingProfileId int `json:"chargingProfileId"`
}

// This field definition of the PullDynamicScheduleUpdate response payload, sent by the Charging Station to the CSMS in response to a PullDynamicScheduleUpdateRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type PullDynamicScheduleUpdateResponse struct {
	ScheduleUpdate *types.ChargingScheduleUpdateType `json:"scheduleUpdate,omitempty" validate:"omitempty"`
	Status         ChargingProfileStatus             `json:"status" validate:"required,chargingProfileStatus201"`
	StatusInfo     *types.StatusInfo                 `json:"statusInfo,omitempty" validate:"omitempty"`
}

// The CSMS may influence the charging power or current drawn from a specific EVSE or
// the entire Charging Station, over a period of time.
// For this purpose, the CSMS calculates a ChargingSchedule to stay within certain limits, then sends a
// PullDynamicScheduleUpdateRequest to the Charging Station. The charging schedule limits may be imposed by any
// external system. The Charging Station responds to this request with a PullDynamicScheduleUpdateResponse.
//
// While charging, the EVSE will continuously adapt the maximum current/power according to the installed
// charging profiles.
type PullDynamicScheduleUpdateFeature struct{}

func (f PullDynamicScheduleUpdateFeature) GetFeatureName() string {
	return PullDynamicScheduleUpdateFeatureName
}

func (f PullDynamicScheduleUpdateFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(PullDynamicScheduleUpdateRequest{})
}

func (f PullDynamicScheduleUpdateFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(PullDynamicScheduleUpdateResponse{})
}

func (r PullDynamicScheduleUpdateRequest) GetFeatureName() string {
	return PullDynamicScheduleUpdateFeatureName
}

func (c PullDynamicScheduleUpdateResponse) GetFeatureName() string {
	return PullDynamicScheduleUpdateFeatureName
}

// Creates a new PullDynamicScheduleUpdateRequest, containing all required fields. There are no optional fields for this message.
func NewPullDynamicScheduleUpdateRequest(chargingProfileId int) *PullDynamicScheduleUpdateRequest {
	return &PullDynamicScheduleUpdateRequest{
		ChargingProfileId: chargingProfileId,
	}
}

// Creates a new PullDynamicScheduleUpdateResponse, containing all required fields. Optional fields may be set afterwards.
func NewPullDynamicScheduleUpdateResponse(status ChargingProfileStatus) *PullDynamicScheduleUpdateResponse {
	return &PullDynamicScheduleUpdateResponse{Status: status}
}
