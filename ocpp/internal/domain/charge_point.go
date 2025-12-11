package domain

import "time"

type Connector struct {
	Id             int         `json:"id"`
	Status         Status      `json:"status"`
	ErrorCode      ErrorCode   `json:"errorCode"`
	LastMeterValue *MeterValue `json:"lastMeterValue,omitempty"`
}

type Transaction struct {
	Id          int        `json:"id"`
	ConnectorId int        `json:"connectorId"`
	IdTag       string     `json:"idTag"`
	StartTime   time.Time  `json:"startTime"`
	StartMeter  int        `json:"startMeter"`
	EndTime     *time.Time `json:"endTime,omitempty"`
	EndMeter    *int       `json:"endMeter,omitempty"`
}

type ChargePoint struct {
	Id            string        `json:"id"`
	LastHeartbeat time.Time     `json:"lastHeartbeat"`
	Connectors    []Connector   `json:"connectors,omitempty"`
	Transactions  []Transaction `json:"transactions,omitempty"`
}
