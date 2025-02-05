package smartcharging

import (
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
	"gopkg.in/go-playground/validator.v9"
)

// DERControlStatusEnumType
type DERControlStatusEnumType string

const (
	DERControlStatusAccepted     DERControlStatusEnumType = "Accepted"
	DERControlStatusRejected     DERControlStatusEnumType = "Rejected"
	DERControlStatusNotSupported DERControlStatusEnumType = "NotSupported"
	DERControlStatusNotFound     DERControlStatusEnumType = "NotFound"
)

func isValidDERControlStatusType(fl validator.FieldLevel) bool {
	status := DERControlStatusEnumType(fl.Field().String())
	switch status {
	case DERControlStatusAccepted, DERControlStatusRejected, DERControlStatusNotSupported, DERControlStatusNotFound:
		return true
	default:
		return false
	}
}

// DERCurvePointsType
type DERCurvePointsType struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// HysteresisType
type HysteresisType struct {
	HysteresisHigh     *float64 `json:"hysteresisHigh,omitempty" validate:"omitempty"`
	HysteresisLow      *float64 `json:"hysteresisLow,omitempty" validate:"omitempty"`
	HysteresisDelay    *float64 `json:"hysteresisDelay,omitempty" validate:"omitempty"`
	HysteresisGradient *float64 `json:"hysteresisGradient,omitempty" validate:"omitempty"`
}

// ReactivePowerParamsType
type ReactivePowerParamsType struct {
	VRef                       *float64 `json:"vRef,omitempty" validate:"omitempty"`
	AutonomousVRefEnable       bool     `json:"autonomousVRefEnable,omitempty" validate:"omitempty"`
	AutonomousVRefTimeConstant *float64 `json:"autonomousVRefTimeConstant,omitempty" validate:"omitempty"`
}

// PowerDuringCessationEnumType
type PowerDuringCessationEnumType string

const (
	PowerDuringCessationActive   PowerDuringCessationEnumType = "Active"
	PowerDuringCessationReactive PowerDuringCessationEnumType = "Reactive"
)

func isValidPowerDuringCessationType(fl validator.FieldLevel) bool {
	status := PowerDuringCessationEnumType(fl.Field().String())
	switch status {
	case PowerDuringCessationActive, PowerDuringCessationReactive:
		return true
	default:
		return false
	}
}

// VoltageParamsType
type VoltageParamsType struct {
	HV10MinMeanValue     *float64                     `json:"hv10MinMeanValue,omitempty" validate:"omitempty"`
	HV10MinMeanTripDelay *float64                     `json:"hv10MinMeanTripDelay,omitempty" validate:"omitempty"`
	PowerDuringCessation PowerDuringCessationEnumType `json:"powerDuringCessation,omitempty" validate:"omitempty,powerDuringCessationType"`
}

// DERUnitEnumType
type DERUnitEnumType string

const (
	DERUnitNotApplicable DERUnitEnumType = "Not_Applicable"
	DERUnitPctMaxW       DERUnitEnumType = "PctMaxW"
	DERUnitPctMaxVar     DERUnitEnumType = "PctMaxVar"
	DERUnitPctWAvail     DERUnitEnumType = "PctWAvail"
	DERUnitPctVarAvail   DERUnitEnumType = "PctVarAvail"
	DERUnitPctEffectiveV DERUnitEnumType = "PctEffectiveV"
)

func isValidDERUnitType(fl validator.FieldLevel) bool {
	status := DERUnitEnumType(fl.Field().String())
	switch status {
	case DERUnitNotApplicable, DERUnitPctMaxW, DERUnitPctMaxVar, DERUnitPctWAvail, DERUnitPctVarAvail, DERUnitPctEffectiveV:
		return true
	default:
		return false
	}
}

// DERCurveType
type DERCurveType struct {
	CurveData           []DERCurvePointsType     `json:"curveData" validate:"required,min=1,max=10,dive"`
	Hysteresis          *HysteresisType          `json:"hysteresis,omitempty" validate:"omitempty"`
	Priority            int                      `json:"priority" validate:"gte=0"`
	ReactivePowerParams *ReactivePowerParamsType `json:"reactivePowerParams,omitempty" validate:"omitempty"`
	VoltageParams       *VoltageParamsType       `json:"voltageParams,omitempty" validate:"omitempty"`
	YUnit               *DERUnitEnumType         `json:"yUnit" validate:"required,derUnitType"`
	ResponseTime        *float64                 `json:"responseTime,omitempty" validate:"omitempty"`
	StartTime           *types.DateTime          `json:"startTime,omitempty" validate:"omitempty"`
	Duration            *float64                 `json:"duration,omitempty" validate:"omitempty"`
}

// DERCurveGetType
type DERCurveGetType struct {
	Curve        DERCurveType             `json:"curve" validate:"required"`
	Id           string                   `json:"id" validate:"required,max=36"`
	CurveType    types.DERControlEnumType `json:"curveType" validate:"required"`
	IsDefault    bool                     `json:"isDefault"`
	IsSuperseded bool                     `json:"isSuperseded"`
}

// EnterServiceType
type EnterServiceType struct {
	Priority    int      `json:"priority" validate:"gte=0"`
	HighVoltage float64  `json:"highVoltage"`
	LowVoltage  float64  `json:"lowVoltage"`
	HighFreq    float64  `json:"highFreq"`
	LowFreq     float64  `json:"lowFreq"`
	Delay       *float64 `json:"delay,omitempty" validate:"omitempty"`
	RampRate    *float64 `json:"rampRate,omitempty" validate:"omitempty"`
}

// EnterServiceGetType
type EnterServiceGetType struct {
	EnterService EnterServiceType `json:"enterService" validate:"required"`
	Id           string           `json:"id" validate:"required,max=36"`
}

// FixedPFType
type FixedPFType struct {
	Priority     int             `json:"priority" validate:"gte=0"`
	Displacement float64         `json:"displacement"`
	Excitation   bool            `json:"excitation"`
	StartTime    *types.DateTime `json:"startTime,omitempty" validate:"omitempty"`
	Duration     *float64        `json:"duration,omitempty" validate:"omitempty"`
}

// FixedPFGetType
type FixedPFGetType struct {
	FixedPF      FixedPFType `json:"fixedPF" validate:"required"`
	Id           string      `json:"id" validate:"required,max=36"`
	IsDefault    bool        `json:"isDefault"`
	IsSuperseded bool        `json:"isSuperseded"`
}

// FixedVarType
type FixedVarType struct {
	Priority  int             `json:"priority" validate:"gte=0"`
	Setpoint  float64         `json:"setpoint"`
	Unit      DERUnitEnumType `json:"unit" validate:"required"`
	StartTime *types.DateTime `json:"startTime,omitempty" validate:"omitempty"`
	Duration  *float64        `json:"duration,omitempty" validate:"omitempty"`
}

// FixedVarGetType
type FixedVarGetType struct {
	FixedVar     FixedVarType `json:"fixedVar" validate:"required"`
	Id           string       `json:"id" validate:"required,max=36"`
	IsDefault    bool         `json:"isDefault"`
	IsSuperseded bool         `json:"isSuperseded"`
}

// FreqDroopType
type FreqDroopType struct {
	Priority     int             `json:"priority" validate:"gte=0"`
	OverFreq     float64         `json:"overFreq"`
	UnderFreq    float64         `json:"underFreq"`
	OverDroop    float64         `json:"overDroop"`
	UnderDroop   float64         `json:"underDroop"`
	ResponseTime float64         `json:"responseTime"`
	StartTime    *types.DateTime `json:"startTime,omitempty" validate:"omitempty"`
	Duration     *float64        `json:"duration,omitempty" validate:"omitempty"`
}

// FreqDroopGetType
type FreqDroopGetType struct {
	FreqDroop    FreqDroopType `json:"freqDroop" validate:"required"`
	Id           string        `json:"id" validate:"required,max=36"`
	IsDefault    bool          `json:"isDefault"`
	IsSuperseded bool          `json:"isSuperseded"`
}

// GradientType
type GradientType struct {
	Priority     int     `json:"priority" validate:"gte=0"`
	Gradient     float64 `json:"gradient"`
	SoftGradient float64 `json:"softGradient"`
}

// GradientGetType
type GradientGetType struct {
	Gradient GradientType `json:"gradient" validate:"required"`
	Id       string       `json:"id" validate:"required,max=36"`
}

// LimitMaxDischargeType
type LimitMaxDischargeType struct {
	Priority                int             `json:"priority" validate:"gte=0"`
	PCTMaxDischargePower    *float64        `json:"pctMaxDischargePower,omitempty" validate:"omitempty"`
	PowerMonitoringMustTrip *DERCurveType   `json:"powerMonitoringMustTrip,omitempty" validate:"omitempty"`
	StartTime               *types.DateTime `json:"startTime,omitempty" validate:"omitempty"`
	Duration                *float64        `json:"duration,omitempty" validate:"omitempty"`
}

// LimitMaxDischargeGetType
type LimitMaxDischargeGetType struct {
	Id                string                `json:"id" validate:"required,max=36"`
	IsDefault         bool                  `json:"isDefault"`
	IsSuperseded      bool                  `json:"isSuperseded"`
	LimitMaxDischarge LimitMaxDischargeType `json:"limitMaxDischarge" validate:"required"`
}

// DERChargingParametersType
type DERChargingParametersType struct {
	EVSupportedDERControl                  []types.DERControlEnumType         `json:"evSupportedDERControl,omitempty" validate:"omitempty,min=1,dive,derControlType"`
	EVOverExcitedMaxDischargePower         *float64                           `json:"evOverExcitedMaxDischargePower,omitempty" validate:"omitempty"`
	EVOverExcitedPowerFactor               *float64                           `json:"evOverExcitedPowerFactor,omitempty" validate:"omitempty"`
	EVUnderExcitedMaxDischargePower        *float64                           `json:"evUnderExcitedMaxDischargePower,omitempty" validate:"omitempty"`
	EVUnderExcitedPowerFactor              *float64                           `json:"evUnderExcitedPowerFactor,omitempty" validate:"omitempty"`
	MaxApparentPower                       *float64                           `json:"maxApparentPower,omitempty" validate:"omitempty"`
	MaxChargeApparentPower                 *float64                           `json:"maxChargeApparentPower,omitempty" validate:"omitempty"`
	MaxChargeApparentPowerL2               *float64                           `json:"maxChargeApparentPower_L2,omitempty" validate:"omitempty"`
	MaxChargeApparentPowerL3               *float64                           `json:"maxChargeApparentPower_L3,omitempty" validate:"omitempty"`
	MaxDischargeApparentPower              *float64                           `json:"maxDischargeApparentPower,omitempty" validate:"omitempty"`
	MaxDischargeApparentPowerL2            *float64                           `json:"maxDischargeApparentPower_L2,omitempty" validate:"omitempty"`
	MaxDischargeApparentPowerL3            *float64                           `json:"maxDischargeApparentPower_L3,omitempty" validate:"omitempty"`
	MaxChargeReactivePower                 *float64                           `json:"maxChargeReactivePower,omitempty" validate:"omitempty"`
	MaxChargeReactivePowerL2               *float64                           `json:"maxChargeReactivePower_L2,omitempty" validate:"omitempty"`
	MaxChargeReactivePowerL3               *float64                           `json:"maxChargeReactivePower_L3,omitempty" validate:"omitempty"`
	MinChargeReactivePower                 *float64                           `json:"minChargeReactivePower,omitempty" validate:"omitempty"`
	MinChargeReactivePowerL2               *float64                           `json:"minChargeReactivePower_L2,omitempty" validate:"omitempty"`
	MinChargeReactivePowerL3               *float64                           `json:"minChargeReactivePower_L3,omitempty" validate:"omitempty"`
	MaxDischargeReactivePower              *float64                           `json:"maxDischargeReactivePower,omitempty" validate:"omitempty"`
	MaxDischargeReactivePowerL2            *float64                           `json:"maxDischargeReactivePower_L2,omitempty" validate:"omitempty"`
	MaxDischargeReactivePowerL3            *float64                           `json:"maxDischargeReactivePower_L3,omitempty" validate:"omitempty"`
	MinDischargeReactivePower              *float64                           `json:"minDischargeReactivePower,omitempty" validate:"omitempty"`
	MinDischargeReactivePowerL2            *float64                           `json:"minDischargeReactivePower_L2,omitempty" validate:"omitempty"`
	MinDischargeReactivePowerL3            *float64                           `json:"minDischargeReactivePower_L3,omitempty" validate:"omitempty"`
	NominalVoltage                         *float64                           `json:"nominalVoltage,omitempty" validate:"omitempty"`
	NominalVoltageOffset                   *float64                           `json:"nominalVoltageOffset,omitempty" validate:"omitempty"`
	MaxNominalVoltage                      *float64                           `json:"maxNominalVoltage,omitempty" validate:"omitempty"`
	MinNominalVoltage                      *float64                           `json:"minNominalVoltage,omitempty" validate:"omitempty"`
	EVInverterManufacturer                 string                             `json:"evInverterManufacturer,omitempty" validate:"omitempty,max=50"`
	EVInverterModel                        string                             `json:"evInverterModel,omitempty" validate:"omitempty,max=50"`
	EVInverterSerialNumber                 string                             `json:"evInverterSerialNumber,omitempty" validate:"omitempty,max=50"`
	EVInverterSwVersion                    string                             `json:"evInverterSwVersion,omitempty" validate:"omitempty,max=50"`
	EVInverterHwVersion                    string                             `json:"evInverterHwVersion,omitempty" validate:"omitempty,max=50"`
	EVIslandingDetectionMethod             []types.IslandingDetectionEnumType `json:"evIslandingDetectionMethod,omitempty" validate:"omitempty,min=1,dive,islandingDetectionType"`
	EVIslandingTripTime                    *float64                           `json:"evIslandingTripTime,omitempty" validate:"omitempty"`
	EVMaximumLevel1DCInjection             *float64                           `json:"evMaximumLevel1DCInjection,omitempty" validate:"omitempty"`
	EVDurationLevel1DCInjection            *float64                           `json:"evDurationLevel1DCInjection,omitempty" validate:"omitempty"`
	EVMaximumLevel2DCInjection             *float64                           `json:"evMaximumLevel2DCInjection,omitempty" validate:"omitempty"`
	EVDurationLevel2DCInjection            *float64                           `json:"evDurationLevel2DCInjection,omitempty" validate:"omitempty"`
	EVReactiveSusceptance                  *float64                           `json:"evReactiveSusceptance,omitempty" validate:"omitempty"`
	EVSessionTotalDischargeEnergyAvailable *float64                           `json:"evSessionTotalDischargeEnergyAvailable,omitempty" validate:"omitempty"`
}

// EVPriceRuleType
type EVPriceRuleType struct {
	EnergyFee       float64 `json:"energyFee"`
	PowerRangeStart float64 `json:"powerRangeStart"`
}

// EVAbsolutePriceScheduleEntryType
type EVAbsolutePriceScheduleEntryType struct {
	Duration    int               `json:"duration"`
	EVPriceRule []EVPriceRuleType `json:"evPriceRule" validate:"required,min=1,max=8,dive"`
}

// EVAbsolutePriceScheduleType
type EVAbsolutePriceScheduleType struct {
	TimeAnchor                     *types.DateTime                    `json:"timeAnchor" validate:"required"`
	Currency                       string                             `json:"currency" validate:"required,len=3"`
	EVAbsolutePriceScheduleEntries []EVAbsolutePriceScheduleEntryType `json:"evAbsolutePriceScheduleEntries" validate:"required,min=1,max=1024,dive"`
	PriceAlgorithm                 string                             `json:"priceAlgorithm" validate:"required,max=2000"`
}

// EVPowerScheduleEntryType
type EVPowerScheduleEntryType struct {
	Duration int     `json:"duration"`
	Power    float64 `json:"power"`
}

// EVPowerScheduleType
type EVPowerScheduleType struct {
	EVPowerScheduleEntries []EVPowerScheduleEntryType `json:"evPowerScheduleEntries" validate:"required,min=1,max=1024,dive"`
	TimeAnchor             *types.DateTime            `json:"timeAnchor" validate:"required"`
}

// EVEnergyOfferType
type EVEnergyOfferType struct {
	EVAbsolutePriceSchedule *EVAbsolutePriceScheduleType `json:"evAbsolutePriceSchedule,omitempty" validate:"omitempty"`
	EVPowerSchedule         *EVPowerScheduleType         `json:"evPowerSchedule" validate:"required"`
}

// V2XChargingParametersType
type V2XChargingParametersType struct {
	MinChargePower        *float64 `json:"minChargePower,omitempty" validate:"omitempty"`
	MinChargePowerL2      *float64 `json:"minChargePower_L2,omitempty" validate:"omitempty"`
	MinChargePowerL3      *float64 `json:"minChargePower_L3,omitempty" validate:"omitempty"`
	MaxChargePower        *float64 `json:"maxChargePower,omitempty" validate:"omitempty"`
	MaxChargePowerL2      *float64 `json:"maxChargePower_L2,omitempty" validate:"omitempty"`
	MaxChargePowerL3      *float64 `json:"maxChargePower_L3,omitempty" validate:"omitempty"`
	MinDischargePower     *float64 `json:"minDischargePower,omitempty" validate:"omitempty"`
	MinDischargePowerL2   *float64 `json:"minDischargePower_L2,omitempty" validate:"omitempty"`
	MinDischargePowerL3   *float64 `json:"minDischargePower_L3,omitempty" validate:"omitempty"`
	MaxDischargePower     *float64 `json:"maxDischargePower,omitempty" validate:"omitempty"`
	MaxDischargePowerL2   *float64 `json:"maxDischargePower_L2,omitempty" validate:"omitempty"`
	MaxDischargePowerL3   *float64 `json:"maxDischargePower_L3,omitempty" validate:"omitempty"`
	MinChargeCurrent      *float64 `json:"minChargeCurrent,omitempty" validate:"omitempty"`
	MaxChargeCurrent      *float64 `json:"maxChargeCurrent,omitempty" validate:"omitempty"`
	MinDischargeCurrent   *float64 `json:"minDischargeCurrent,omitempty" validate:"omitempty"`
	MaxDischargeCurrent   *float64 `json:"maxDischargeCurrent,omitempty" validate:"omitempty"`
	MinVoltage            *float64 `json:"minVoltage,omitempty" validate:"omitempty"`
	MaxVoltage            *float64 `json:"maxVoltage,omitempty" validate:"omitempty"`
	EVTargetEnergyRequest *float64 `json:"evTargetEnergyRequest,omitempty" validate:"omitempty"`
	EVMinEnergyRequest    *float64 `json:"evMinEnergyRequest,omitempty" validate:"omitempty"`
	EVMaxEnergyRequest    *float64 `json:"evMaxEnergyRequest,omitempty" validate:"omitempty"`
	EVMinV2XEnergyRequest *float64 `json:"evMinV2XEnergyRequest,omitempty" validate:"omitempty"`
	EVMaxV2XEnergyRequest *float64 `json:"evMaxV2XEnergyRequest,omitempty" validate:"omitempty"`
	TargetSoC             int      `json:"targetSoC,omitempty" validate:"omitempty,gte=0,lte=100"`
}

func init() {
	_ = types.Validate.RegisterValidation("derControlStatusType", isValidDERControlStatusType)
	_ = types.Validate.RegisterValidation("powerDuringCessationType", isValidPowerDuringCessationType)
	_ = types.Validate.RegisterValidation("derUnitType", isValidDERUnitType)
}
