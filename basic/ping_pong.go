package main

import (
	"fmt"
	"time"
)

func RunPingPong() {
	fmt.Println("Starting game....")
	pingPong(3)
	fmt.Println("======= Game over =======")
}

func pingPong(numPings int) {
	pings := make(chan struct{}, numPings)
	pongs := make(chan struct{}, numPings)

	go pinger(pings, numPings)
	go ponger(pings, pongs)
	go func() {
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("Pong done")
	}()

	<-pings
	<-pongs
}

func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("Sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("Pings done")
	close(pongs)
}
