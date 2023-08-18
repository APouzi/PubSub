package main

type SubPub struct {
	Subscribers map[string][]func(string)
}

func NewPubSub() *SubPub {
	return &SubPub{
		Subscribers: make(map[string][]func(string)),
	}
}

func (pb *SubPub) Subscribe(topic string, fn func(string)) {
	pb.Subscribers[topic] = append(pb.Subscribers[topic], fn)
}

func (pb *SubPub) Publish(topic string, message string) {
	for _, fnHnd := range pb.Subscribers[topic] {
		fnHnd(message)
	}
}
