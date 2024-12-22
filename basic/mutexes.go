package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

func writeLoop(m map[int]int, mu *sync.Mutex) {

	for i := 0; i < 100; i++ {
		mu.Lock()
		m[i] = i
		mu.Unlock()
	}

}

func readLoop(m map[int]int, mu *sync.Mutex) {
	for {
		mu.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mu.Unlock()
	}
}

func RunReadWrite() {
	m := map[int]int{}
	mu := &sync.Mutex{}
	go writeLoop(m, mu)
	go readLoop(m, mu)

	// Stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}
