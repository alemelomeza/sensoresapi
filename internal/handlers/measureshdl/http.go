package measureshdl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alemelomeza/sensor-api/internal/core/domain"

	"github.com/alemelomeza/sensor-api/internal/core/ports"
)

// HTTPHandler ...
type HTTPHandler struct {
	measureService ports.MeasuresService
}

// NewHTTPHandler ...
func NewHTTPHandler(measuresService ports.MeasuresService) *HTTPHandler {
	return &HTTPHandler{
		measureService: measuresService,
	}
}

// Create ...
func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	var measure domain.Measure

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&measure); err != nil {
		http.Error(w, fmt.Sprint("Datos incorrectos"), http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Registro exitoso de la medición"))
	w.WriteHeader(http.StatusOK)
}
