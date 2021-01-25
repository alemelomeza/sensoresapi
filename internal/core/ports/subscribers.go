package ports

import (
	"github.com/alemelomeza/sensor-api/internal/core/domain"
)

// SubscribersRepository ...
type SubscribersRepository interface {
	Save(subscribe domain.Subscriber) error
	Get() (domain.Subscriber, error)
}

// SubscribersService ...
type SubscribersService interface {
	Subscribe(subscribe domain.Subscriber) error
	Get() ([]domain.Subscriber, error)
}
