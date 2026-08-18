package worker

type NotificationAttempt struct{ Acknowledged bool }

func (a *NotificationAttempt) Run(deliver func() error) (err error) {
	defer func() {
		if err == nil {
			a.Acknowledged = true
		}
	}()
	return deliver()
}
