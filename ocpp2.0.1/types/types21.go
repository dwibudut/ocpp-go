package types

import (
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
	"gopkg.in/go-playground/validator.v9"
)

const (
	V21Subprotocol = "ocpp2.1"
)

// EnergyTransferModeEnumType
type EnergyTransferModeEnumType string

const (
	ACSinglePhase EnergyTransferModeEnumType = "AC_single_phase"
	ACTwoPhase    EnergyTransferModeEnumType = "AC_two_phase"
	ACThreePhase  EnergyTransferModeEnumType = "AC_three_phase"
	ACBPT         EnergyTransferModeEnumType = "AC_BPT"
	ACBPTDER      EnergyTransferModeEnumType = "AC_BPT_DER"
	ACDER         EnergyTransferModeEnumType = "AC_DER"
	DCBPT         EnergyTransferModeEnumType = "DC_BPT"
	DCACDP        EnergyTransferModeEnumType = "DC_ACDP"
	DCACDPBPT     EnergyTransferModeEnumType = "DC_ACDP_BPT"
	WPT           EnergyTransferModeEnumType = "WPT"
)

func isValidEnergyTransferModeType(fl validator.FieldLevel) bool {
	status := EnergyTransferModeEnumType(fl.Field().String())
	switch status {
	case ACSinglePhase, ACTwoPhase, ACThreePhase, ACBPT, ACBPTDER, ACDER, DCBPT, DCACDP, DCACDPBPT, WPT:
		return true
	default:
		return false
	}
}

const (
	Monday    DayOfWeekEnumType = "Monday"
	Tuesday   DayOfWeekEnumType = "Tuesday"
	Wednesday DayOfWeekEnumType = "Wednesday"
	Thursday  DayOfWeekEnumType = "Thursday"
	Friday    DayOfWeekEnumType = "Friday"
	Saturday  DayOfWeekEnumType = "Saturday"
	Sunday    DayOfWeekEnumType = "Sunday"
)

func isValidDayOfWeekType(fl validator.FieldLevel) bool {
	purposeType := DayOfWeekEnumType(fl.Field().String())
	switch purposeType {
	case Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday:
		return true
	default:
		return false
	}
}

// DayOfWeekEnumType
type DayOfWeekEnumType string

const (
	AC EvseKindEnumType = "AC"
	DC EvseKindEnumType = "DC"
)

func isValidEvseKindType(fl validator.FieldLevel) bool {
	purposeType := EvseKindEnumType(fl.Field().String())
	switch purposeType {
	case AC, DC:
		return true
	default:
		return false
	}
}

// EvseKindEnumType
type EvseKindEnumType string

// TariffConditionsType
type TariffConditionsType struct {
	StartTimeOfDay  string              `json:"startTimeOfDay,omitempty" validate:"omitempty"`
	EndTimeOfDay    string              `json:"endTimeOfDay,omitempty" validate:"omitempty"`
	DayOfWeek       []DayOfWeekEnumType `json:"dayOfWeek,omitempty" validate:"omitempty,min=1,max=7,dive,dayOfWeek"`
	ValidFromDate   string              `json:"validFromDate,omitempty" validate:"omitempty"`
	ValidToDate     string              `json:"validToDate,omitempty" validate:"omitempty"`
	EvseKind        EvseKindEnumType    `json:"evseKind,omitempty" validate:"omitempty,evseKind"`
	MinEnergy       *float64            `json:"minEnergy,omitempty" validate:"omitempty"`
	MaxEnergy       *float64            `json:"maxEnergy,omitempty" validate:"omitempty"`
	MinCurrent      *float64            `json:"minCurrent,omitempty" validate:"omitempty"`
	MaxCurrent      *float64            `json:"maxCurrent,omitempty" validate:"omitempty"`
	MinPower        *float64            `json:"minPower,omitempty" validate:"omitempty"`
	MaxPower        *float64            `json:"maxPower,omitempty" validate:"omitempty"`
	MinTime         int                 `json:"minTime,omitempty" validate:"omitempty"`
	MaxTime         int                 `json:"maxTime,omitempty" validate:"omitempty"`
	MinChargingTime int                 `json:"minChargingTime,omitempty" validate:"omitempty"`
	MaxChargingTime int                 `json:"maxChargingTime,omitempty" validate:"omitempty"`
	MinIdleTime     int                 `json:"minIdleTime,omitempty" validate:"omitempty"`
	MaxIdleTime     int                 `json:"maxIdleTime,omitempty" validate:"omitempty"`
}

// TariffEnergyPriceType
type TariffEnergyPriceType struct {
	PriceKwh   int `json:"priceKwh"`
	Conditions int `json:"conditions,omitempty" validate:"omitempty,min=0"`
}

// TaxRateType
type TaxRateType struct {
	Type  string  `json:"type" validate:"required,max=20"`
	Tax   float64 `json:"tax"`
	Stack int     `json:"stack,omitempty" validate:"omitempty,gte=0"`
}

// TariffEnergyType
type TariffEnergyType struct {
	Prices   []TariffEnergyPriceType `json:"prices" validate:"required,min=1,dive"`
	TaxRates []TaxRateType           `json:"taxRates,omitempty" validate:"omitempty,min=1,max=5,dive"`
}

// MessageContentType
type MessageContentType struct {
	Format   MessageFormatType `json:"format" validate:"required"`
	Language string            `json:"language,omitempty" validate:"omitempty,max=8"`
	Content  string            `json:"content" validate:"required,max=1024"`
}

// TariffTimePriceType
type TariffTimePriceType struct {
	PriceMinute float64               `json:"priceMinute"`
	Conditions  *TariffConditionsType `json:"conditions,omitempty" validate:"omitempty"`
}

// TariffTimeType
type TariffTimeType struct {
	Prices   []TariffTimePriceType `json:"prices" validate:"required,min=1,dive"`
	TaxRates []TaxRateType         `json:"taxRates,omitempty" validate:"omitempty,min=1,max=5,dive"`
}

// TariffConditionsFixedType
type TariffConditionsFixedType struct {
	StartTimeOfDay     string              `json:"startTimeOfDay,omitempty" validate:"omitempty"`
	EndTimeOfDay       string              `json:"endTimeOfDay,omitempty" validate:"omitempty"`
	DayOfWeek          []DayOfWeekEnumType `json:"dayOfWeek,omitempty" validate:"omitempty,min=1,max=7,dive,dayOfWeek"`
	ValidFromDate      string              `json:"validFromDate,omitempty" validate:"omitempty"`
	ValidToDate        string              `json:"validToDate,omitempty" validate:"omitempty"`
	EvseKind           EvseKindEnumType    `json:"evseKind,omitempty" validate:"omitempty,evseKind"`
	PaymentBrand       string              `json:"paymentBrand,omitempty" validate:"omitempty,max=20"`
	PaymentRecognition string              `json:"paymentRecognition,omitempty" validate:"omitempty,max=20"`
}

// TariffFixedPriceType
type TariffFixedPriceType struct {
	Conditions *TariffConditionsFixedType `json:"conditions,omitempty" validate:"omitempty"`
	PriceFixed float64                    `json:"priceFixed"`
}

// TariffFixedType
type TariffFixedType struct {
	Prices   []TariffFixedPriceType `json:"prices" validate:"required,min=1,dive"`
	TaxRates []TaxRateType          `json:"taxRates,omitempty" validate:"omitempty,min=1,max=5,dive"`
}

// PriceType
type PriceType struct {
	ExclTax  *float64      `json:"exclTax,omitempty" validate:"omitempty"`
	InclTax  *float64      `json:"inclTax,omitempty" validate:"omitempty"`
	TaxRates []TaxRateType `json:"taxRates,omitempty" validate:"omitempty,min=1,max=5,dive"`
}

// TariffType
type TariffType struct {
	TariffId         string               `json:"tariffId,required" validate:"required,max=60"`
	Description      []MessageContentType `json:"description,omitempty" validate:"omitempty,min=1,max=10,dive"`
	Currency         string               `json:"currency" validate:"required,len=3"`
	Energy           *TariffEnergyType    `json:"energy,omitempty" validate:"omitempty"`
	ValidFrom        *DateTime            `json:"validFrom,omitempty" validate:"omitempty"`
	ChargingTime     *TariffTimeType      `json:"chargingTime,omitempty" validate:"omitempty"`
	IdleTime         *TariffTimeType      `json:"idleTime,omitempty" validate:"omitempty"`
	FixedFee         *TariffFixedType     `json:"fixedFee,omitempty" validate:"omitempty"`
	ReservationTime  *TariffTimeType      `json:"reservationTime,omitempty" validate:"omitempty"`
	ReservationFixed *TariffFixedType     `json:"reservationFixed,omitempty" validate:"omitempty"`
	MinCost          *PriceType           `json:"minCost,omitempty" validate:"omitempty"`
	MaxCost          *PriceType           `json:"maxCost,omitempty" validate:"omitempty"`
}

type CostDimensionEnumType string

const (
	CostDimensionEnergy       CostDimensionEnumType = "Energy"
	CostDimensionMaxCurrent   CostDimensionEnumType = "MaxCurrent"
	CostDimensionMinCurrent   CostDimensionEnumType = "MinCurrent"
	CostDimensionMaxPower     CostDimensionEnumType = "MaxPower"
	CostDimensionMinPower     CostDimensionEnumType = "MinPower"
	CostDimensionIdleTIme     CostDimensionEnumType = "IdleTIme"
	CostDimensionChargingTime CostDimensionEnumType = "ChargingTime"
)

func isValidCostDimensionType(fl validator.FieldLevel) bool {
	chargingLimitSource := CostDimensionEnumType(fl.Field().String())
	switch chargingLimitSource {
	case CostDimensionEnergy, CostDimensionMaxCurrent, CostDimensionMinCurrent, CostDimensionMaxPower, CostDimensionMinPower, CostDimensionIdleTIme, CostDimensionChargingTime:
		return true
	default:
		return false
	}
}

// CostDimensionType
type CostDimensionType struct {
	Type   CostDimensionEnumType `json:"type" validate:"required,costDimensionType"`
	Volume float64               `json:"volume"`
}

// ChargingPeriodType
type ChargingPeriodType struct {
	Dimensions  []CostDimensionType `json:"dimensions,omitempty" validate:"omitempty,min=1,dive"`
	TariffId    string              `json:"tariffId,omitempty" validate:"omitempty,max=60"`
	StartPeriod DateTime            `json:"startPeriod" validate:"required"`
}

// TariffCostEnumType
type TariffCostEnumType string

const (
	TariffNormalCost TariffCostEnumType = "NormalCost"
	TariffMinCost    TariffCostEnumType = "MinCost"
	TariffMaxCost    TariffCostEnumType = "MaxCost"
)

func isValidTariffCostType(fl validator.FieldLevel) bool {
	chargingLimitSource := TariffCostEnumType(fl.Field().String())
	switch chargingLimitSource {
	case TariffNormalCost, TariffMinCost, TariffMaxCost:
		return true
	default:
		return false
	}
}

// TotalPriceType
type TotalPriceType struct {
	ExclTax *float64 `json:"exclTax,omitempty" validate:"omitempty"`
	InclTax *float64 `json:"inclTax,omitempty" validate:"omitempty"`
}

// TotalCostType
type TotalCostType struct {
	Currency         string             `json:"currency" validate:"required,len=3"`
	TypeOfCost       TariffCostEnumType `json:"typeOfCost" validate:"required,tariffCostType"`
	Fixed            *PriceType         `json:"fixed,omitempty" validate:"omitempty"`
	Energy           *PriceType         `json:"energy,omitempty" validate:"omitempty"`
	ChargingTime     *PriceType         `json:"chargingTime,omitempty" validate:"omitempty"`
	IdleTime         *PriceType         `json:"idleTime,omitempty" validate:"omitempty"`
	ReservationTime  *PriceType         `json:"reservationTime,omitempty" validate:"omitempty"`
	ReservationFixed *PriceType         `json:"reservationFixed,omitempty" validate:"omitempty"`
	Total            TotalPriceType     `json:"total" validate:"required"`
}

// TotalUsageType
type TotalUsageType struct {
	Energy          float64 `json:"energy"`
	ChargingTime    int     `json:"chargingTime"`
	IdleTime        int     `json:"idleTime"`
	ReservationTime int     `json:"reservationTime,omitempty" validate:"omitempty"`
}

// CostDetailsType
type CostDetailsType struct {
	ChargingPeriods    []ChargingPeriodType `json:"chargingPeriods,omitempty" validate:"omitempty,min=1,dive"`
	TotalCost          TotalCostType        `json:"totalCost" validate:"required"`
	TotalUsage         TotalUsageType       `json:"totalUsage" validate:"required"`
	FailureToCalculate bool                 `json:"failureToCalculate,omitempty" validate:"omitempty"`
	FailureReason      string               `json:"failureReason,omitempty" validate:"omitempty,max=500"`
}

// PreconditioningStatusEnumType
type PreconditioningStatusEnumType string

const (
	PreconditioningStatusUnknown         PreconditioningStatusEnumType = "Unknown"
	PreconditioningStatusReady           PreconditioningStatusEnumType = "Ready"
	PreconditioningStatusNotReady        PreconditioningStatusEnumType = "NotReady"
	PreconditioningStatusPreconditioning PreconditioningStatusEnumType = "Preconditioning"
)

func isValidPreconditioningStatusType(fl validator.FieldLevel) bool {
	status := PreconditioningStatusEnumType(fl.Field().String())
	switch status {
	case PreconditioningStatusUnknown, PreconditioningStatusReady, PreconditioningStatusNotReady, PreconditioningStatusPreconditioning:
		return true
	default:
		return false
	}
}

// TransactionLimitType
type TransactionLimitType struct {
	MaxCost   *float64 `json:"maxCost,omitempty" validate:"omitempty"`
	MaxEnergy *float64 `json:"maxEnergy,omitempty" validate:"omitempty"`
	MaxTime   int      `json:"maxTime,omitempty" validate:"omitempty"`
	MaxSoC    int      `json:"maxSoC,omitempty" validate:"omitempty,min=0,max=100"`
}

// SigningMethodEnumStringType
type SigningMethodEnumStringType string

const (
	SigningMethodECDSASecp192k1SHA256      SigningMethodEnumStringType = "ECDSA-secp192k1-SHA256"
	SigningMethodECDSASecp256k1SHA256      SigningMethodEnumStringType = "ECDSA-secp256k1-SHA256"
	SigningMethodECDSASecp192r1SHA256      SigningMethodEnumStringType = "ECDSA-secp192r1-SHA256"
	SigningMethodECDSASecp256r1SHA256      SigningMethodEnumStringType = "ECDSA-secp256r1-SHA256"
	SigningMethodECDSABrainpool256r1SHA256 SigningMethodEnumStringType = "ECDSA-brainpool256r1-SHA256"
	SigningMethodECDSASecp384r1SHA256      SigningMethodEnumStringType = "ECDSA-secp384r1-SHA256"
	SigningMethodECDSABrainpool384r1SHA256 SigningMethodEnumStringType = "ECDSA-brainpool384r1-SHA256"
)

// OperationModeEnumType
type OperationModeEnumType string

const (
	OperationModeIdle               OperationModeEnumType = "Idle"
	OperationModeChargingOnly       OperationModeEnumType = "ChargingOnly"
	OperationModeCentralSetpoint    OperationModeEnumType = "CentralSetpoint"
	OperationModeExternalSetpoint   OperationModeEnumType = "ExternalSetpoint"
	OperationModeExternalLimits     OperationModeEnumType = "ExternalLimits"
	OperationModeCentralFrequency   OperationModeEnumType = "CentralFrequency"
	OperationModeLocalFrequency     OperationModeEnumType = "LocalFrequency"
	OperationModeLocalLoadBalancing OperationModeEnumType = "LocalLoadBalancing"
)

func isOperationModeType(fl validator.FieldLevel) bool {
	status := OperationModeEnumType(fl.Field().String())
	switch status {
	case OperationModeIdle, OperationModeChargingOnly, OperationModeCentralSetpoint, OperationModeExternalSetpoint, OperationModeExternalLimits, OperationModeCentralFrequency, OperationModeLocalFrequency, OperationModeLocalLoadBalancing:
		return true
	default:
		return false
	}
}

// V2XFreqWattPointType
type V2XFreqWattPointType struct {
	Frequency float64 `json:"frequency"`
	Power     float64 `json:"power"`
}

// V2XSignalWattPointType
type V2XSignalWattPointType struct {
	Signal float64 `json:"signal"`
	Power  float64 `json:"power"`
}

// LimitAtSoCType
type LimitAtSoCType struct {
	Soc   int     `json:"soc" validate:"gte=0,lte=100"`
	Limit float64 `json:"limit"`
}

// RationalNumberType
type RationalNumberType struct {
	Exponent int `json:"exponent"`
	Value    int `json:"value"`
}

// PriceRuleType
type PriceRuleType struct {
	ParkingFeePeriod              int                 `json:"parkingFeePeriod,omitempty" validate:"omitempty"`
	CarbonDioxideEmission         int                 `json:"carbonDioxideEmission,omitempty" validate:"omitempty,gte=0"`
	RenewableGenerationPercentage int                 `json:"renewableGenerationPercentage,omitempty" validate:"omitempty,gte=0,lte=100"`
	EnergyFee                     RationalNumberType  `json:"energyFee" validate:"required"`
	ParkingFee                    *RationalNumberType `json:"parkingFee,omitempty" validate:"omitempty"`
	PowerRangeStart               RationalNumberType  `json:"powerRangeStart" validate:"required"`
}

// PriceRuleStackType
type PriceRuleStackType struct {
	Duration  int             `json:"duration"`
	PriceRule []PriceRuleType `json:"priceRule" validate:"required,min=1,max=8,dive"`
}

// TaxRuleType
type TaxRuleType struct {
	TaxRuleID                   int                `json:"taxRuleID" validate:"gte=0"`
	TaxRuleName                 string             `json:"taxRuleName,omitempty" validate:"omitempty,max=100"`
	TaxIncludedInPrice          bool               `json:"taxIncludedInPrice,omitempty" validate:"omitempty"`
	AppliesToEnergyFee          bool               `json:"appliesToEnergyFee"`
	AppliesToParkingFee         bool               `json:"appliesToParkingFee"`
	AppliesToOverstayFee        bool               `json:"appliesToOverstayFee"`
	AppliesToMinimumMaximumCost bool               `json:"appliesToMinimumMaximumCost"`
	TaxRate                     RationalNumberType `json:"taxRate" validate:"required"`
}

// OverstayRuleType
type OverstayRuleType struct {
	OverstayFee             RationalNumberType `json:"overstayFee" validate:"required"`
	OverstayRuleDescription string             `json:"overstayRuleDescription,omitempty" validate:"omitempty,max=32"`
	StartTime               int                `json:"startTime"`
	OverstayFeePeriod       int                `json:"overstayFeePeriod"`
}

// OverstayRuleListType
type OverstayRuleListType struct {
	OverstayPowerThreshold *RationalNumberType `json:"overstayPowerThreshold,omitempty" validate:"omitempty"`
	OverstayRule           []OverstayRuleType  `json:"overstayRule" validate:"required,min=1,max=5,dive"`
	OverstayTimeThreshold  int                 `json:"overstayTimeThreshold,omitempty" validate:"omitempty"`
}

// AdditionalSelectedServicesType
type AdditionalSelectedServicesType struct {
	ServiceFee  RationalNumberType `json:"serviceFee" validate:"required"`
	ServiceName string             `json:"serviceName" validate:"required,max=90"`
}

// AbsolutePriceScheduleType
type AbsolutePriceScheduleType struct {
	TimeAnchor                 DateTime                         `json:"timeAnchor" validate:"required"`
	PriceScheduleID            int                              `json:"priceScheduleID" validate:"gte=0"`
	PriceScheduleDescription   string                           `json:"priceScheduleDescription,omitempty" validate:"omitempty,max=160"`
	Currency                   string                           `json:"currency" validate:"required,len=3"`
	Language                   string                           `json:"language" validate:"required,max=8"`
	PriceAlgorithm             string                           `json:"priceAlgorithm" validate:"required,max=2000"`
	MinimumCost                *RationalNumberType              `json:"minimumCost,omitempty" validate:"omitempty"`
	MaximumCost                *RationalNumberType              `json:"maximumCost,omitempty" validate:"omitempty"`
	PriceRuleStacks            []PriceRuleStackType             `json:"priceRuleStacks" validate:"required,min=1,max=1024,dive"`
	TaxRules                   []TaxRuleType                    `json:"taxRules,omitempty" validate:"omitempty,min=1,max=10,dive"`
	OverstayRuleList           *OverstayRuleListType            `json:"overstayRuleList,omitempty" validate:"omitempty"`
	AdditionalSelectedServices []AdditionalSelectedServicesType `json:"additionalSelectedServices,omitempty" validate:"omitempty,min=1,max=5,dive"`
}

// PriceLevelScheduleEntryType
type PriceLevelScheduleEntryType struct {
	Duration   int `json:"duration"`
	PriceLevel int `json:"priceLevel" validate:"gte=0"`
}

// PriceLevelScheduleType
type PriceLevelScheduleType struct {
	PriceLevelScheduleEntries []PriceLevelScheduleEntryType `json:"priceLevelScheduleEntries" validate:"required,min=1,max=100,dive"`
	TimeAnchor                DateTime                      `json:"timeAnchor" validate:"required"`
	PriceScheduleId           int                           `json:"priceScheduleId" validate:"gte=0"`
	PriceScheduleDescription  string                        `json:"priceScheduleDescription,omitempty" validate:"omitempty,max=32"`
	NumberOfPriceLevels       int                           `json:"numberOfPriceLevels" validate:"gte=0"`
}

// ChargingScheduleUpdateType
type ChargingScheduleUpdateType struct {
	Limit              *float64 `json:"limit,omitempty" validate:"omitempty"`
	LimitL2            *float64 `json:"limit_L2,omitempty" validate:"omitempty"`
	LimitL3            *float64 `json:"limit_L3,omitempty" validate:"omitempty"`
	DischargeLimit     int      `json:"dischargeLimit,omitempty" validate:"omitempty,lte=0"`
	DischargeLimitL2   int      `json:"dischargeLimit_L2,omitempty" validate:"omitempty,lte=0"`
	DischargeLimitL3   int      `json:"dischargeLimit_L3,omitempty" validate:"omitempty,lte=0"`
	Setpoint           int      `json:"setpoint,omitempty" validate:"omitempty"`
	SetpointL2         int      `json:"setpoint_L2,omitempty" validate:"omitempty"`
	SetpointL3         int      `json:"setpoint_L3,omitempty" validate:"omitempty"`
	SetpointReactive   int      `json:"setpointReactive,omitempty" validate:"omitempty"`
	SetpointReactiveL2 int      `json:"setpointReactive_L2,omitempty" validate:"omitempty"`
	SetpointReactiveL3 int      `json:"setpointReactive_L3,omitempty" validate:"omitempty"`
}

// AddressType
type AddressType struct {
	Name       string `json:"name" validate:"required,max=50"`
	Address1   string `json:"address1" validate:"required,max=100"`
	Address2   string `json:"address2,omitempty" validate:"omitempty,max=100"`
	City       string `json:"city" validate:"required,max=100"`
	PostalCode string `json:"postalCode,omitempty" validate:"omitempty,max=20"`
	Country    string `json:"country" validate:"required,max=50"`
}

// DERControlEnumType
type DERControlEnumType string

const (
	DERControlEnterService            DERControlEnumType = "EnterService"
	DERControlFreqDroop               DERControlEnumType = "FreqDroop"
	DERControlFreqWatt                DERControlEnumType = "FreqWatt"
	DERControlFixedPFAbsorb           DERControlEnumType = "FixedPFAbsorb"
	DERControlFixedPFInject           DERControlEnumType = "FixedPFInject"
	DERControlFixedVar                DERControlEnumType = "FixedVar"
	DERControlGradients               DERControlEnumType = "Gradients"
	DERControlHFMustTrip              DERControlEnumType = "HFMustTrip"
	DERControlHFMayTrip               DERControlEnumType = "HFMayTrip"
	DERControlHVMustTrip              DERControlEnumType = "HVMustTrip"
	DERControlHVMomCess               DERControlEnumType = "HVMomCess"
	DERControlHVMayTrip               DERControlEnumType = "HVMayTrip"
	DERControlLimitMaxDischarge       DERControlEnumType = "LimitMaxDischarge"
	DERControlLFMustTrip              DERControlEnumType = "LFMustTrip"
	DERControlLVMustTrip              DERControlEnumType = "LVMustTrip"
	DERControlLVMomCess               DERControlEnumType = "LVMomCess"
	DERControlLVMayTrip               DERControlEnumType = "LVMayTrip"
	DERControlPowerMonitoringMustTrip DERControlEnumType = "PowerMonitoringMustTrip"
	DERControlVoltVar                 DERControlEnumType = "VoltVar"
	DERControlVoltWatt                DERControlEnumType = "VoltWatt"
	DERControlWattPF                  DERControlEnumType = "WattPF"
	DERControlWattVar                 DERControlEnumType = "WattVar"
)

func isValidDERControlType(fl validator.FieldLevel) bool {
	status := DERControlEnumType(fl.Field().String())
	switch status {
	case DERControlEnterService, DERControlFreqDroop, DERControlFreqWatt, DERControlFixedPFAbsorb, DERControlFixedPFInject, DERControlFixedVar, DERControlGradients, DERControlHFMustTrip, DERControlHFMayTrip, DERControlHVMustTrip, DERControlHVMomCess, DERControlHVMayTrip, DERControlLimitMaxDischarge, DERControlLFMustTrip, DERControlLVMustTrip, DERControlLVMomCess, DERControlLVMayTrip, DERControlPowerMonitoringMustTrip, DERControlVoltVar, DERControlVoltWatt, DERControlWattPF, DERControlWattVar:
		return true
	default:
		return false
	}
}

// IslandingDetectionEnumType
type IslandingDetectionEnumType string

const (
	IslandingDetectionNoAntiIslandingSupport IslandingDetectionEnumType = "NoAntiIslandingSupport"
	IslandingDetectionRoCoF                  IslandingDetectionEnumType = "RoCoF"
	IslandingDetectionUVP_OVP                IslandingDetectionEnumType = "UVP_OVP"
	IslandingDetectionUFP_OFP                IslandingDetectionEnumType = "UFP_OFP"
	IslandingDetectionVoltageVectorShift     IslandingDetectionEnumType = "VoltageVectorShift"
	IslandingDetectionZeroCrossingDetection  IslandingDetectionEnumType = "ZeroCrossingDetection"
	IslandingDetectionOtherPassive           IslandingDetectionEnumType = "OtherPassive"
	IslandingDetectionImpedanceMeasurement   IslandingDetectionEnumType = "ImpedanceMeasurement"
	IslandingDetectionImpedanceAtFrequency   IslandingDetectionEnumType = "ImpedanceAtFrequency"
	IslandingDetectionSlipModeFrequencyShift IslandingDetectionEnumType = "SlipModeFrequencyShift"
	IslandingDetectionSandiaFrequencyShift   IslandingDetectionEnumType = "SandiaFrequencyShift"
	IslandingDetectionSandiaVoltageShift     IslandingDetectionEnumType = "SandiaVoltageShift"
	IslandingDetectionFrequencyJump          IslandingDetectionEnumType = "FrequencyJump"
	IslandingDetectionRCLQFactor             IslandingDetectionEnumType = "RCLQFactor"
	IslandingDetectionOtherActive            IslandingDetectionEnumType = "OtherActive"
)

func isValidIslandingDetectionType(fl validator.FieldLevel) bool {
	status := IslandingDetectionEnumType(fl.Field().String())
	switch status {
	case IslandingDetectionNoAntiIslandingSupport, IslandingDetectionRoCoF, IslandingDetectionUVP_OVP, IslandingDetectionUFP_OFP, IslandingDetectionVoltageVectorShift, IslandingDetectionZeroCrossingDetection, IslandingDetectionOtherPassive, IslandingDetectionImpedanceMeasurement, IslandingDetectionImpedanceAtFrequency, IslandingDetectionSlipModeFrequencyShift, IslandingDetectionSandiaFrequencyShift, IslandingDetectionSandiaVoltageShift, IslandingDetectionFrequencyJump, IslandingDetectionRCLQFactor, IslandingDetectionOtherActive:
		return true
	default:
		return false
	}
}

// ControlModeEnumType
type ControlModeEnumType string

const (
	ControlModeScheduledControl ControlModeEnumType = "ScheduledControl"
	ControlModeDynamicControl   ControlModeEnumType = "DynamicControl"
)

func isValidControlModeType(fl validator.FieldLevel) bool {
	status := ControlModeEnumType(fl.Field().String())
	switch status {
	case ControlModeScheduledControl, ControlModeDynamicControl:
		return true
	default:
		return false
	}
}

// GridEventFaultEnumType
type GridEventFaultEnumType string

const (
	GridEventFaultCurrentImbalance GridEventFaultEnumType = "CurrentImbalance"
	GridEventFaultLocalEmergency   GridEventFaultEnumType = "LocalEmergency"
	GridEventFaultLowInputPower    GridEventFaultEnumType = "LowInputPower"
	GridEventFaultOverCurrent      GridEventFaultEnumType = "OverCurrent"
	GridEventFaultOverFrequency    GridEventFaultEnumType = "OverFrequency"
	GridEventFaultOverVoltage      GridEventFaultEnumType = "OverVoltage"
	GridEventFaultPhaseRotation    GridEventFaultEnumType = "PhaseRotation"
	GridEventFaultRemoteEmergency  GridEventFaultEnumType = "RemoteEmergency"
	GridEventFaultUnderFrequency   GridEventFaultEnumType = "UnderFrequency"
	GridEventFaultUnderVoltage     GridEventFaultEnumType = "UnderVoltage"
	GridEventFaultVoltageImbalance GridEventFaultEnumType = "VoltageImbalance"
)

func isValidGridEventFaultType(fl validator.FieldLevel) bool {
	status := GridEventFaultEnumType(fl.Field().String())
	switch status {
	case GridEventFaultCurrentImbalance, GridEventFaultLocalEmergency, GridEventFaultLowInputPower, GridEventFaultOverCurrent, GridEventFaultOverFrequency, GridEventFaultOverVoltage, GridEventFaultPhaseRotation, GridEventFaultRemoteEmergency, GridEventFaultUnderFrequency, GridEventFaultUnderVoltage, GridEventFaultVoltageImbalance:
		return true
	default:
		return false
	}
}

// PaymentStatusEnumType
type PaymentStatusEnumType string

const (
	PaymentStatusSettled  PaymentStatusEnumType = "Settled"
	PaymentStatusCanceled PaymentStatusEnumType = "Canceled"
	PaymentStatusRejected PaymentStatusEnumType = "Rejected"
	PaymentStatusFailed   PaymentStatusEnumType = "Failed"
)

func isValidPaymentStatusType(fl validator.FieldLevel) bool {
	status := PaymentStatusEnumType(fl.Field().String())
	switch status {
	case PaymentStatusSettled, PaymentStatusCanceled, PaymentStatusRejected, PaymentStatusFailed:
		return true
	default:
		return false
	}
}

// CertificateStatusSourceEnum
type CertificateStatusSourceEnumType string

const (
	CertificateStatusSourceCRL  CertificateStatusSourceEnumType = "CRL"
	CertificateStatusSourceOCSP CertificateStatusSourceEnumType = "OCSP"
)

func isValidCertificateStatusSource(fl validator.FieldLevel) bool {
	status := CertificateStatusSourceEnumType(fl.Field().String())
	switch status {
	case CertificateStatusSourceCRL, CertificateStatusSourceOCSP:
		return true
	default:
		return false
	}
}

// CertificateStatusEnumType
type CertificateStatusEnumType string

const (
	CertificateStatusGood    CertificateStatusEnumType = "Good"
	CertificateStatusRevoked CertificateStatusEnumType = "Revoked"
	CertificateStatusUnknown CertificateStatusEnumType = "Unknown"
	CertificateStatusFailed  CertificateStatusEnumType = "Failed"
)

func isCertificateStatus(fl validator.FieldLevel) bool {
	status := CertificateStatusEnumType(fl.Field().String())
	switch status {
	case CertificateStatusGood, CertificateStatusRevoked, CertificateStatusUnknown, CertificateStatusFailed:
		return true
	default:
		return false
	}
}

// CertificateStatusType
type CertificateStatusType struct {
	CertificateHashData CertificateHashData             `json:"certificateHashData" validate:"required"`
	Source              CertificateStatusSourceEnumType `json:"source" validate:"required"`
	Status              CertificateStatusEnumType       `json:"status" validate:"required,certificateStatus"`
	NextUpdate          *types.DateTime                 `json:"nextUpdate" validate:"required"`
}

func init() {
	_ = Validate.RegisterValidation("energyTransferModeType", isValidEnergyTransferModeType)
	_ = Validate.RegisterValidation("dayOfWeekType", isValidDayOfWeekType)
	_ = Validate.RegisterValidation("evseKindType", isValidEvseKindType)
	_ = Validate.RegisterValidation("costDimensionType", isValidCostDimensionType)
	_ = Validate.RegisterValidation("tariffCostType", isValidTariffCostType)
	_ = Validate.RegisterValidation("preconditioningStatusType", isValidPreconditioningStatusType)
	_ = Validate.RegisterValidation("operationModeType", isOperationModeType)
	_ = Validate.RegisterValidation("derControlType", isValidDERControlType)
	_ = Validate.RegisterValidation("islandingDetectionType", isValidIslandingDetectionType)
	_ = Validate.RegisterValidation("controlModeType", isValidControlModeType)
	_ = Validate.RegisterValidation("gridEventFaultType", isValidGridEventFaultType)
	_ = Validate.RegisterValidation("paymentStatusType", isValidPaymentStatusType)
	_ = Validate.RegisterValidation("certificateStatusSource", isValidCertificateStatusSource)
	_ = Validate.RegisterValidation("certificateStatus", isCertificateStatus)
}
