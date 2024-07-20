package events

import "errors"

type EventHandlerDispatcher struct {
	handlers map[string][]EventHandler
}

var EventHandlerAlreadExists = errors.New("Invalid Handler, its already exists.")

func NewEventDispatcher() *EventHandlerDispatcher {
	return &EventHandlerDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (ed *EventHandlerDispatcher) Register(eventName string, eventHandler EventHandler) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == eventHandler {
				return EventHandlerAlreadExists
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], eventHandler)
	return nil
}
