package ports

import (
	"github.com/alemelomeza/sensor-api/internal/core/domain"
)

// MeasuresService ...
type MeasuresService interface {
	Receive(measure domain.Measure) error
}
