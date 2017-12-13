package state

import (
	"reflect"
	"sync"
)

// Op is the operation perfored on the state object.
type Op string

const (
	// SET a new key/value is added.
	SET Op = "SET"

	// DELETE a key is removed.
	DELETE Op = "DELETE"
)

// Listerner is a callback fired when the state change.
type Listerner func(ops Op, key interface{}, state *State)

// State stores key/value pairs.
type State struct {
	m         *sync.Map
	listeners map[interface{}][]Listerner
	global    Listerner
}

func New() *State {
	return &State{
		m:         &sync.Map{},
		listeners: make(map[interface{}][]Listerner),
	}
}

// Set adds key/value to the state store. If there was any previous value it
// will be overwritten.
//
//Nothing happens if the previous value and the  new value are equal.
func (s *State) Set(key, value interface{}) {
	if v, ok := s.m.Load(key); ok {
		if reflect.DeepEqual(v, value) {
			return
		}
	}
	s.m.Store(key, value)
	if ls, ok := s.listeners[key]; ok {
		for _, l := range ls {
			l(SET, key, s)
		}
	}
	if s.global != nil {
		s.global(SET, key, s)
	}
}

// Get returns value stored by key.
func (s *State) Get(key interface{}) (interface{}, bool) {
	return s.m.Load(key)
}

// GetOk returns value stored by key and boolean indicating if the key was
// present in the store.
func (s *State) GetOk(key interface{}) (interface{}, bool) {
	return s.m.Load(key)
}

// Listen register f as a callback which will be called whenever value
// associated with key changes.
func (s *State) Listen(key interface{}, f Listerner) {
	if ls, ok := s.listeners[key]; ok {
		ls = append(ls, f)
		s.listeners[key] = ls
	} else {
		s.listeners[key] = []Listerner{f}
	}
}

// SetGlobal registers f as a global listener for the state object.
func (s *State) SetGlobal(l Listerner) {
	s.global = l
}
