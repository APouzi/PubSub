package main

import "sync"

type PubSubConCur struct {
	mutex sync.RWMutex
	Subscribers map[string][]chan string
}
func NewPubSubConCur() *PubSubConCur{
	return &PubSubConCur{
		mutex: sync.RWMutex{},
		Subscribers: make(map[string][]chan string),
	}
}
func(pb *PubSubConCur) Subscribe(topic string) chan string{
	pb.mutex.Lock()
	ch := make(chan string)
	pb.Subscribers[topic] = append(pb.Subscribers[topic], ch)
	pb.mutex.Unlock()
	return ch
}

func (pb *PubSubConCur) Publish(topic, msg string) {
	pb.mutex.RLock()
	for _, channel := range pb.Subscribers[topic]{
		channel <- msg
	}
	pb.mutex.RUnlock()
}

