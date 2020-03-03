package cqrs

type AggregateStore interface {
	EntityStore
	EventSink
}

type EntityStore interface {
	Load() (interface{}, error)
}

type EventSink interface {
	AppendEvents([]Event) error
}
