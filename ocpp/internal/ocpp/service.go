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
	
}

func NewOcppService() *OcppService {
	return &OcppService{}
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

	fmt.Printf("[%s] Parsed OCPP: %+v\n", cpId, msg.Action)

	switch msg.Action {
	case "BootNotification":
		response, err := handlers.HandleBootNotification(msg)
		if err != nil {
			return err
		}

		responseBytes, err := domain.NewCallResult(msg.Id, response)
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if err := conn.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
			return fmt.Errorf("failed to write response: %w", err)
		}
	}

	return nil
}
