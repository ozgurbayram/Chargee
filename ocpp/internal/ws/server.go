package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) WsHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	parts := strings.Split(path, "/")

	cpId := ""
	if len(parts) >= 3 {
		cpId = parts[2]
	}

	responseHeader := http.Header{}
	if r.Header.Get("Sec-WebSocket-Protocol") != "" {
		responseHeader.Set("Sec-WebSocket-Protocol", r.Header.Get("Sec-WebSocket-Protocol"))
	}

	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		fmt.Println("Error while establish websocket connection:", err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("Error while reading ws message:", err)
			break
		}

		if err := s.handleIncomingOcppMessage(message, cpId); err != nil {
			fmt.Println("Error while writing ws message:", err)
			break
		}
	}

}

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

func (s *Server) handleIncomingOcppMessage(message []byte, cpId string) error {
	var msg OcppMessage

	if err := json.Unmarshal(message, &msg); err != nil {
		return fmt.Errorf("invalid ocpp message format: %w", err)
	}

	fmt.Printf("[%s] Parsed OCPP: %+v\n", cpId, msg.Action)
	return nil
}
