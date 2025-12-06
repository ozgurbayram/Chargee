package domain

import "time"

type ChargePoint struct {
	Id            string    `json:"id"`
	LastHeartbeat time.Time `json:"lastHeartbeat"`
}
