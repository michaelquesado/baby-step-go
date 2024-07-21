package events

import (
	"sync"
	"time"
)

type Event interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandler interface {
	Handler(e Event, wg *sync.WaitGroup)
}

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Dispatcher(e Event) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear() error
}
