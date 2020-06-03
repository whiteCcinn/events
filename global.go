package events

import "reflect"

type Event interface {
	On(eventName string, task interface{}) error
	Fire(eventName string, params ...interface{}) ([]reflect.Value, error)
	FireBackground(eventName string, params ...interface{}) (chan []reflect.Value, error)
	Clear(eventName string) error
	ClearEvents()
	HasEvent(eventName string) bool
	Events() []string
	EventCount() int
}

var globalEvent = New()

// Default global trigger options.
func On(eventName string, task interface{}) error {
	return globalEvent.On(eventName, task)
}

func Fire(eventName string, params ...interface{}) ([]reflect.Value, error) {
	return globalEvent.Fire(eventName, params...)
}

func FireBackground(eventName string, params ...interface{}) (chan []reflect.Value, error) {
	return globalEvent.FireBackground(eventName, params...)
}

func Clear(eventName string) error {
	return globalEvent.Clear(eventName)
}

func ClearEvents() {
	globalEvent.ClearEvents()
}

func HasEvent(eventName string) bool {
	return globalEvent.HasEvent(eventName)
}

func Events() []string {
	return globalEvent.Events()
}

func EventCount() int {
	return globalEvent.EventCount()
}
