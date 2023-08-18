package main

import "sync"

type PubSubConCur2 struct {
	Mutex       sync.RWMutex
	Subscribers map[string]chan string
}

func NewPubSubConCur2() *PubSubConCur2 {
	return &PubSubConCur2{
		Mutex:       sync.RWMutex{},
		Subscribers: make(map[string]chan string),
	}
}

func(pb *PubSubConCur2) Subscribe(topic string) <-chan string{
	pb.Mutex.Lock()
	ch := make(chan string)
	pb.Subscribers[topic] = ch
	pb.Mutex.Unlock()
	return ch
}