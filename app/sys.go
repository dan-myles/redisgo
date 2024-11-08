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

func (s *sys) Get(key string) string {
	return s.kv[key]
}

func (s *sys) Set(key string, val string) {
	s.kv[key] = val
}
