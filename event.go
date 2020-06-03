package events

import (
	"errors"
	"reflect"
	"sync"
)

func New() Event {
	return new(event)
}

type event struct {
	functionMap *sync.Map
}

func (e *event) On(eventName string, task interface{}) error {

	if _, ok := e.functionMap.Load(eventName); ok {
		return errors.New("eventName already defined")
	}

	if reflect.ValueOf(task).Type().Kind() != reflect.Func {
		return errors.New("task is not a function")
	}

	e.functionMap.Store(eventName, task)

	return nil
}

func (e *event) Fire(eventName string, params ...interface{}) ([]reflect.Value, error) {
	f, in, err := e.read(eventName, params...)
	if err != nil {
		return nil, err
	}
	result := f.Call(in)

	return result, nil
}

func (e *event) FireBackground(eventName string, params ...interface{}) (chan []reflect.Value, error) {
	f, in, err := e.read(eventName, params...)
	if err != nil {
		return nil, err
	}

	results := make(chan []reflect.Value)
	go func() {
		results <- f.Call(in)
	}()

	return results, nil
}

func (e *event) Clear(eventName string) error {
	if _, ok := e.functionMap.Load(eventName); !ok {
		return errors.New("event not defined")
	}
	e.functionMap.Delete(eventName)

	return nil
}

func (e *event) ClearEvents() {
	e.functionMap = new(sync.Map)
}

func (e *event) HasEvent(eventName string) bool {
	_, ok := e.functionMap.Load(eventName)

	return ok
}

func (e *event) Events() []string {
	events := make([]string, 0)

	e.functionMap.Range(func(k, v interface{}) bool {
		events = append(events, k.(string))
		return true
	})

	return events
}

func (e *event) EventCount() int {
	i := 0

	e.functionMap.Range(func(k, v interface{}) bool {
		i++
		return true
	})

	return i
}

func (e *event) read(eventName string, params ...interface{}) (reflect.Value, []reflect.Value, error) {
	task, ok := e.functionMap.Load(eventName)
	if !ok {
		return reflect.Value{}, nil, errors.New("no task found for event")
	}

	f := reflect.ValueOf(task)
	if len(params) != f.Type().NumIn() {
		return reflect.Value{}, nil, errors.New("parameter mismatched")
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	return f, in, nil
}
