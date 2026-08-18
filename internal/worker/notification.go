package worker

import "github.com/kekelele996/subsea-cable-repair-control-service/internal/service"

func DeliverNotifications(s *service.NotificationService, batch string, messages []string, send func([]string) error) error {
	return s.Deliver(batch, messages, send)
}
