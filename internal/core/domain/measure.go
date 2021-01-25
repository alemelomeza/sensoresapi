package domain

import (
	"time"
)

// Measure ...
type Measure struct {
	SensorID  string    `json:"sensorId"`
	Timestamp time.Time `json:"timestamp"`
	Value     int       `json:"Value"`
}
