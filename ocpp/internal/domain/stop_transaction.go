package domain

type StopTransactionRequest struct {
	IdTag           string       `json:"idTag,omitempty"`
	MeterStop       int          `json:"meterStop"`
	Timestamp       string       `json:"timestamp"`
	TransactionId   int          `json:"transactionId"`
	Reason          Reason       `json:"reason,omitempty"`
	TransactionData []MeterValue `json:"transactionData,omitempty"`
}

type Reason string

const (
	ReasonEmergencyStop  Reason = "EmergencyStop"
	ReasonEVDisconnected Reason = "EVDisconnected"
	ReasonHardReset      Reason = "HardReset"
	ReasonLocal          Reason = "Local"
	ReasonOther          Reason = "Other"
	ReasonPowerLoss      Reason = "PowerLoss"
	ReasonReboot         Reason = "Reboot"
	ReasonRemote         Reason = "Remote"
	ReasonSoftReset      Reason = "SoftReset"
	ReasonUnlockCommand  Reason = "UnlockCommand"
	ReasonDeAuthorized   Reason = "DeAuthorized"
)

type StopTransactionResponse struct {
	IdTagInfo *IdTagInfo `json:"idTagInfo,omitempty"`
}
