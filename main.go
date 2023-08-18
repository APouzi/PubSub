package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	// Sequintial Version
	PB := NewPubSub()
	PB.Subscribe("tech", func(msg string) {
		fmt.Println("sub1 recieved message:" + msg)
	})
	PB.Subscribe("cars", func(msg string) {
		fmt.Println("sub2 recieved message:" + msg)
	})
	PB.Publish("tech","vr is cool!")
	PB.Publish("cars","they get you goin")


	// Concurrent version
	PBC := NewPubSubConCur()
	var wg sync.WaitGroup
	wg.Add(2)
	sub1 := PBC.Subscribe("tech")
	sub2 := PBC.Subscribe("cars")
	ctx1, _ := context.WithTimeout(context.Background(),time.Second * 5)
	ctx2, _ := context.WithTimeout(context.Background(),time.Second * 5)
	go func(ctx context.Context){
		for {
			select{
				case v1 := <- sub1:
					fmt.Println("sub1 concurrent message:",v1)
				case <-ctx.Done():
					wg.Done()
					return
				}
		}
		
	}(ctx1)
	
	go func(ctx context.Context){
		for {
			select{
				case v2 := <- sub2:
					fmt.Println("sub2 concurrent message:",v2)
				case <-ctx.Done():
					wg.Done()
					return
			}
		}
	}(ctx2)
	
	
	PBC.Publish("tech", "vr is cool!")
	PBC.Publish("cars", "cars are cool!")
	


	//Concurrent FanOut Version:
	PBC2 := NewPubSubConCur2()
	wg.Add(2)
	subFan1 := PBC2.Subscribe("tech")
	subFan2ChanCopy := subFan1
	
	go func(ctx context.Context){
		for {
			select{
				case v1 := <- subFan1:
					fmt.Println("subFan1 concurrent message:",v1)
				case <-ctx.Done():
					wg.Done()
					return
				}
		}
		
	}(ctx1)
	
	go func(ctx context.Context){
		for {
			select{
				case v2 := <- subFan2ChanCopy:
					fmt.Println("subFan2 concurrent message:",v2)
				case <-ctx.Done():
					wg.Done()
					return
			}
		}
	}(ctx2)
	
}