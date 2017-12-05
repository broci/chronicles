package state

import "sync"

type State struct {
	m         *sync.Map
	listeners map[interface{}][]func(interface{}, *State)
}

func New() *State {
	return &State{
		m:         &sync.Map{},
		listeners: make(map[interface{}][]func(interface{}, *State)),
	}
}
func (s *State) Set(key, value interface{}) {
	s.m.Store(key, value)
	if ls, ok := s.listeners[key]; ok {
		for _, l := range ls {
			l(key, s)
		}
	}
}
func (s *State) Get(key interface{}) (interface{}, bool) {
	return s.m.Load(key)
}

func (s *State) Listen(key interface{}, f func(key interface{}, state *State)) {
	if ls, ok := s.listeners[key]; ok {
		ls = append(ls, f)
		s.listeners[key] = ls
	} else {
		s.listeners[key] = []func(interface{}, *State){f}
	}
}
