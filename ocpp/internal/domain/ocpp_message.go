package domain

import (
	"encoding/json"
	"fmt"
)

type OcppMessage struct {
	Type    int
	Id      string
	Action  string
	Message json.RawMessage
}

func (m *OcppMessage) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 4 {
		return fmt.Errorf("invalid OCPP message length: expected 4, got %d", len(raw))
	}

	if err := json.Unmarshal(raw[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(raw[1], &m.Id); err != nil {
		return err
	}
	if err := json.Unmarshal(raw[2], &m.Action); err != nil {
		return err
	}

	m.Message = raw[3]
	return nil
}
