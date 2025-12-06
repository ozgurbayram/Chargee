package domain

type HeartbeatRequest struct {
}

type HeartbeatResponse struct {
	CurrentTime string `json:"currentTime"`
}
