package sse

import (
	"encoding/json"
	"fmt"
	"net/http"

	serializer "github.com/mernat/sso-clean-arch/api/json"
	"github.com/mernat/sso-clean-arch/models"
	sseUseCase "github.com/mernat/sso-clean-arch/usecase/sse"
)

type eventServiceHandler struct {
	service sseUseCase.Service
}

func NewSSEServiceHandler(s sseUseCase.Service) *eventServiceHandler {
	return &eventServiceHandler{
		service: s,
	}
}

func (f *eventServiceHandler) Stream(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	//Get broker
	broker := f.service.GetBroker()

	// Signal the broker that we have a new connection
	broker.NewClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		broker.ClosingClients <- messageChan
	}()

	go func() {
		// Listen to connection close and un-register messageChan
		<-r.Context().Done()
		broker.ClosingClients <- messageChan
	}()

	for {
		// Write to the ResponseWriter
		// Server Sent Events compatible
		fmt.Fprintf(w, "data: %s\n\n", <-messageChan)

		// Flush the data immediatly instead of buffering it for later.
		flusher.Flush()
	}
}

func (f *eventServiceHandler) BroadcastMessage(w http.ResponseWriter, r *http.Request) {
	var message models.MessageBroker

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "cant parse request",
		})
		return
	}

	byteData, err := json.Marshal(message)

	if err != nil {
		serializer.JSON(w, http.StatusBadRequest, &serializer.GenericResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "parging struct to byte",
		})
		return
	}

	broker := f.service.GetBroker()

	broker.Notifier <- byteData

	serializer.JSON(w, http.StatusOK, &serializer.GenericResponse{
		Success: true,
		Code:    http.StatusOK,
	})

}
