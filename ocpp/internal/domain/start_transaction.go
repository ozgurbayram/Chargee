package domain

type StartTransactionRequest struct {
	ConnectorId   int    `json:"connectorId"`
	IdTag         string `json:"idTag"`
	MeterStart    int    `json:"meterStart"`
	ReservationId int    `json:"reservationId,omitempty"`
	Timestamp     string `json:"timestamp"`
}

type StartTransactionResponse struct {
	IdTagInfo     IdTagInfo `json:"idTagInfo"`
	TransactionId int       `json:"transactionId"`
}
