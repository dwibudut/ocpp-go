package smartcharging

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Set Charging Profile (CSMS -> CS) --------------------

const UpdateDynamicScheduleFeatureName = "UpdateDynamicSchedule"

// The field definition of the UpdateDynamicSchedule request payload sent by the CSMS to the Charging Station.
type UpdateDynamicScheduleRequest struct {
	ChargingProfileId int                              `json:"chargingProfileId"`
	ScheduleUpdate    types.ChargingScheduleUpdateType `json:"scheduleUpdate" validate:"required"`
}

// This field definition of the UpdateDynamicSchedule response payload, sent by the Charging Station to the CSMS in response to a UpdateDynamicScheduleRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type UpdateDynamicScheduleResponse struct {
	Status     ChargingProfileStatus `json:"status" validate:"required,chargingProfileStatus201"`
	StatusInfo *types.StatusInfo     `json:"statusInfo,omitempty" validate:"omitempty"`
}

type UpdateDynamicScheduleFeature struct{}

func (f UpdateDynamicScheduleFeature) GetFeatureName() string {
	return UpdateDynamicScheduleFeatureName
}

func (f UpdateDynamicScheduleFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(UpdateDynamicScheduleRequest{})
}

func (f UpdateDynamicScheduleFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(UpdateDynamicScheduleResponse{})
}

func (r UpdateDynamicScheduleRequest) GetFeatureName() string {
	return UpdateDynamicScheduleFeatureName
}

func (c UpdateDynamicScheduleResponse) GetFeatureName() string {
	return UpdateDynamicScheduleFeatureName
}

// Creates a new UpdateDynamicScheduleRequest, containing all required fields. There are no optional fields for this message.
func NewUpdateDynamicScheduleRequest(chargingProfileId int, scheduleUpdate types.ChargingScheduleUpdateType) *UpdateDynamicScheduleRequest {
	return &UpdateDynamicScheduleRequest{
		ChargingProfileId: chargingProfileId,
		ScheduleUpdate:    scheduleUpdate,
	}
}

// Creates a new UpdateDynamicScheduleResponse, containing all required fields. Optional fields may be set afterwards.
func NewUpdateDynamicScheduleResponse(status ChargingProfileStatus) *UpdateDynamicScheduleResponse {
	return &UpdateDynamicScheduleResponse{Status: status}
}
