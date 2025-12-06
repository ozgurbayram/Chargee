package domain

type StatusNotificationRequest struct {
	ConnectorId     int       `json:"connectorId"`
	ErrorCode       ErrorCode `json:"errorCode"`
	Status          Status    `json:"status"`
	Timestamp       string    `json:"timestamp,omitempty"`
	Info            string    `json:"info,omitempty"`
	VendorId        string    `json:"vendorId,omitempty"`
	VendorErrorCode string    `json:"vendorErrorCode,omitempty"`
}

type ErrorCode string

const (
	ErrorCodeConnectorLockFailure ErrorCode = "ConnectorLockFailure"
	ErrorCodeEVCommunicationError ErrorCode = "EVCommunicationError"
	ErrorCodeGroundFailure        ErrorCode = "GroundFailure"
	ErrorCodeHighTemperature      ErrorCode = "HighTemperature"
	ErrorCodeInternalError        ErrorCode = "InternalError"
	ErrorCodeLocalListConflict    ErrorCode = "LocalListConflict"
	ErrorCodeNoError              ErrorCode = "NoError"
	ErrorCodeOtherError           ErrorCode = "OtherError"
	ErrorCodeOverCurrentFailure   ErrorCode = "OverCurrentFailure"
	ErrorCodePowerMeterFailure    ErrorCode = "PowerMeterFailure"
	ErrorCodePowerSwitchFailure   ErrorCode = "PowerSwitchFailure"
	ErrorCodeReaderFailure        ErrorCode = "ReaderFailure"
	ErrorCodeResetFailure         ErrorCode = "ResetFailure"
	ErrorCodeUnderVoltage         ErrorCode = "UnderVoltage"
	ErrorCodeOverVoltage          ErrorCode = "OverVoltage"
	ErrorCodeWeakSignal           ErrorCode = "WeakSignal"
)

type Status string

const (
	StatusAvailable     Status = "Available"
	StatusPreparing     Status = "Preparing"
	StatusCharging      Status = "Charging"
	StatusSuspendedEVSE Status = "SuspendedEVSE"
	StatusSuspendedEV   Status = "SuspendedEV"
	StatusFinishing     Status = "Finishing"
	StatusReserved      Status = "Reserved"
	StatusUnavailable   Status = "Unavailable"
	StatusFaulted       Status = "Faulted"
)

type StatusNotificationResponse struct {
}
