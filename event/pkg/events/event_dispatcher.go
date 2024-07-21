package events

import (
	"errors"
	"sync"
)

type EventHandlerDispatcher struct {
	handlers map[string][]EventHandler
}

var ErrEventHandlerAlreadExists = errors.New("invalid handler, its already exists")

func NewEventDispatcher() *EventHandlerDispatcher {
	return &EventHandlerDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (ed *EventHandlerDispatcher) Register(eventName string, eventHandler EventHandler) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == eventHandler {
				return ErrEventHandlerAlreadExists
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], eventHandler)
	return nil
}

func (ed *EventHandlerDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandler)
	return nil
}

func (ed *EventHandlerDispatcher) Has(eventName string, eventHandler EventHandler) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == eventHandler {
				return true
			}
		}
	}
	return false
}

func (ed *EventHandlerDispatcher) Dispatch(e Event) error {
	if handlers, ok := ed.handlers[e.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, h := range handlers {
			wg.Add(1)
			h.Handler(e, wg)
		}
		wg.Wait()
	}
	return nil
}
