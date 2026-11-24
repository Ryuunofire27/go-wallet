package events

import "go-wallet/internal/shared/src/domain/events"

type EventCollector interface {
	Add(event ...events.Event)
	Collect() []events.Event
}
