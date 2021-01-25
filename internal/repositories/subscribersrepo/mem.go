package subscribersrepo

import (
	"errors"

	"github.com/alemelomeza/sensor-api/internal/core/domain"
)

// MemRepo ...
type MemRepo struct {
	items []domain.Subscriber
}

// NewRepository ...
func NewRepository() *MemRepo {
	return &MemRepo{
		items: []domain.Subscriber{},
	}
}

// Save ...
func (m *MemRepo) Save(subscribe domain.Subscriber) error {
	for _, s := range m.items {
		if subscribe.Endpoint == s.Endpoint {
			return errors.New("Datos incorrectos")
		}
		m.items = append(m.items, subscribe)
		break
	}
	return nil
}

// Get ...
func (m *MemRepo) Get() ([]domain.Subscriber, error) {
	if len(m.items) == 0 {
		return nil, errors.New("Error interno")
	}
	return m.items, nil
}
