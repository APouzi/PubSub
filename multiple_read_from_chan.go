package main

import "sync"

type PubSubConCur2 struct {
	Mutex       sync.RWMutex
	Subscribers map[string][]chan string
}

func NewPubSubConCur2() *PubSubConCur {
	return &PubSubConCur{
		mutex:       sync.RWMutex{},
		Subscribers: make(map[string][]chan string),
	}
}