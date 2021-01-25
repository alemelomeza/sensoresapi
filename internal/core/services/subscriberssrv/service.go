package subscriberssrv

import (
	"errors"

	"github.com/alemelomeza/sensor-api/internal/core/domain"

	"github.com/alemelomeza/sensor-api/internal/core/ports"
)

// SubscribersService ...
type SubscribersService struct {
	subscribersRepository ports.SubscribersRepository
}

// NewService ...
func NewService(subscribersRepository ports.SubscribersRepository) *SubscribersService {
	return &SubscribersService{
		subscribersRepository: subscribersRepository,
	}
}

// Subscribe ...
func (s *SubscribersService) Subscribe(subscriber domain.Subscriber) error {
	if err := s.subscribersRepository.Save(subscriber); err != nil {
		return errors.New("Error interno")
	}
	return nil
}

// Get ...
func (s *SubscribersService) Get() ([]domain.Subscriber, error) {
	subscribers := s.subscribersRepository.Get()
	if len(subscribers) == 0 {
		return nil, errors.New("Error interno")
	}
	return subscribers, nil
}
