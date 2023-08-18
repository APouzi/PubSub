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

// Subscribe is allowing appending the channels that are "representing" the subscribers that are listening. The locks will mean no other subscriber can write to this. 
// Notice that we are returning a channel reciever, this is for safety and encapsulation, to make sure that it can only recieve values and not send them.
func(pb *PubSubConCur) Subscribe(topic string) <-chan string{
	pb.mutex.Lock()
	ch := make(chan string)
	pb.Subscribers[topic] = append(pb.Subscribers[topic], ch)
	pb.mutex.Unlock()
	return ch
}

// Here we have the Read lock and unlock. This is checking if are writing to the subscribers, if not a single coroutine can read from subscribers, otherwise it will be able to read. 
func (pb *PubSubConCur) Publish(topic, msg string) {
	pb.mutex.RLock()
	for _, channel := range pb.Subscribers[topic]{
		channel <- msg
	}
	pb.mutex.RUnlock()
}