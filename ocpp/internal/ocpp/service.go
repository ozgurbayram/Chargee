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

type MessageHandler func(message domain.OcppMessage, cpID string, repository domain.ChargePointRepository) (*domain.OcppMessage, error)

type Service struct {
	handlers   map[string]MessageHandler
	repository domain.ChargePointRepository
}

func NewOcppService(repository domain.ChargePointRepository) *Service {
	return &Service{
		handlers: map[string]MessageHandler{
			"BootNotification":   handlers.HandleBootNotification,
			"Heartbeat":          handlers.HandleHeartbeat,
			"StatusNotification": handlers.StatusNotificationHandler,
		},
		repository: repository,
	}
}

var upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *Service) WsHandler(w http.ResponseWriter, r *http.Request) {

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

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error while closing websocket connection:", err)
		}
	}(conn)

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

func (s *Service) handleIncomingOcppMessage(message []byte, cpId string, conn *websocket.Conn) error {
	var msg domain.OcppMessage

	if err := json.Unmarshal(message, &msg); err != nil {
		return fmt.Errorf("invalid ocpp message format: %w", err)
	}

	handler, ok := s.handlers[msg.Action]

	if !ok {
		return fmt.Errorf("unknown ocpp message action: %s", msg.Action)
	}

	response, err := handler(msg, cpId, s.repository)
	if err != nil {
		return err
	}

	if err := s.writeResponse(conn, *response); err != nil {
		return err
	}

	return nil
}

func (s *Service) writeResponse(conn *websocket.Conn, response domain.OcppMessage) error {

	responseBytes, err := json.Marshal([]interface{}{
		response.Type,
		response.Id,
		response.Action,
		response.Message,
	})

	if err != nil {
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}
