package subscribershdl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alemelomeza/sensor-api/internal/core/domain"
	"github.com/alemelomeza/sensor-api/internal/core/ports"
)

// HTTPHandler ...
type HTTPHandler struct {
	subscribersService ports.SubscribersService
}

// NewHTTPHandler ...
func NewHTTPHandler(subscribersService ports.SubscribersService) *HTTPHandler {
	return &HTTPHandler{
		subscribersService: subscribersService,
	}
}

// Create ...
func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	var subscriber domain.Subscriber
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&subscriber); err != nil {
		http.Error(w, fmt.Sprint("Datos incorrectos"), http.StatusMethodNotAllowed)
	}
	err := h.subscribersService.Subscribe(subscriber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("Registro exitoso de suscriptor"))
	w.WriteHeader(http.StatusOK)
}
