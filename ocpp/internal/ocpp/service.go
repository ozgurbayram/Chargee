package ocpp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ocpp/internal/domain"
	"ocpp/internal/handlers"
	"strings"

	"github.com/gorilla/websocket"
)

type OcppService struct {
	handlers map[string]OcppMessageHandler
}

type OcppMessageHandler func(message domain.OcppMessage) (*domain.OcppMessage, error)

func NewOcppService() *OcppService {
	return &OcppService{
		handlers: map[string]OcppMessageHandler{
			"BootNotification": handlers.HandleBootNotification,
			"Heartbeat":        handlers.HandleHeartbeat,
		},
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *OcppService) WsHandler(w http.ResponseWriter, r *http.Request) {

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

		if err := s.handleIncomingOcppMessage(message, cpId, conn); err != nil {
			fmt.Println("Error while handling ocpp message:", err)
			break
		}
	}

}

func (s *OcppService) handleIncomingOcppMessage(message []byte, cpId string, conn *websocket.Conn) error {
	var msg domain.OcppMessage

	if err := json.Unmarshal(message, &msg); err != nil {
		return fmt.Errorf("invalid ocpp message format: %w", err)
	}

	handler, ok := s.handlers[msg.Action]

	if !ok {
		return fmt.Errorf("unknown ocpp message action: %s", msg.Action)
	}

	response, err := handler(msg)
	if err != nil {
		return err
	}

	if err := s.writeResponse(conn, *response); err != nil {
		return err
	}

	return nil
}

func (s *OcppService) writeResponse(conn *websocket.Conn, response domain.OcppMessage) error {
	responseBytes, err := domain.NewCallResult(response.Id, response)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}
