package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetData() interface{}
	GetTime() time.Time
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handle EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handle EventHandlerInterface) error
	Has(eventName string) bool
	Clear() error
}
