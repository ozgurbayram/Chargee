package domain

import (
	"encoding/json"
	"fmt"
)

const (
	MessageTypeCall       = 2
	MessageTypeCallResult = 3
	MessageTypeCallError  = 4
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

	if len(raw) < 3 || len(raw) > 5 {
		return fmt.Errorf("invalid OCPP message length: expected 3-5, got %d", len(raw))
	}

	if err := json.Unmarshal(raw[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(raw[1], &m.Id); err != nil {
		return err
	}

	if m.Type == 2 { // CALL
		if len(raw) != 4 {
			return fmt.Errorf("CALL message must have 4 elements, got %d", len(raw))
		}
		if err := json.Unmarshal(raw[2], &m.Action); err != nil {
			return err
		}
		m.Message = raw[3]
	} else if m.Type == 3 { // CALLRESULT
		if len(raw) != 3 {
			return fmt.Errorf("CALLRESULT message must have 3 elements, got %d", len(raw))
		}
		m.Action = ""
		m.Message = raw[2]
	} else if m.Type == 4 { // CALLERROR
		if len(raw) != 5 {
			return fmt.Errorf("CALLERROR message must have 5 elements, got %d", len(raw))
		}
		if err := json.Unmarshal(raw[2], &m.Action); err != nil {
			return err
		}
		m.Message = raw[4] // error details
	} else {
		return fmt.Errorf("unknown OCPP message type: %d", m.Type)
	}

	return nil
}
