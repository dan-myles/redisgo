package main

import "sync"

type sys struct {
	kv  map[string]string
	mtx *sync.Mutex
}

func NewSys() *sys {
	return &sys{
		kv:  make(map[string]string),
		mtx: &sync.Mutex{},
	}
}

func (s *sys) Get(key string) (string, bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	val, ok := s.kv[key]
	return val, ok
}

func (s *sys) Set(key string, val string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.kv[key] = val
}

func (s *sys) Delete(key string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	delete(s.kv, key)
}
