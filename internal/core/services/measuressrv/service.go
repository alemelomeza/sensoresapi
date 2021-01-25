package measuressrv

import (
	"github.com/alemelomeza/sensor-api/internal/core/domain"
)

// MeasuresService ...
type MeasuresService struct{}

// NewService ...
func NewService() *MeasuresService {
	return &MeasuresService{}
}

// Receive ...
func (s *MeasuresService) Receive(measure domain.Measure) error {
	return nil
}
