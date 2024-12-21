package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple send message
func sendEmail(message string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		time.Sleep(time.Microsecond * 250)
		fmt.Printf("Email recived: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string, wg *sync.WaitGroup) {
	sendEmail(message, wg)
	time.Sleep(time.Microsecond * 500)
	fmt.Println("-------------------------------------------")
}

func simulateEmailSend() {
	fmt.Println("===========Email simulation===============")
	var wg sync.WaitGroup
	wg.Add(3)
	test("Email number 1", &wg)
	test("Email number 2", &wg)
	test("Email number 3", &wg)

	wg.Wait()
	fmt.Println("Done sending email")
}

func channels() {
	var wg sync.WaitGroup

	ageChan := make(chan int, 6)
	wg.Add(6)
	for i := 10; i <= 15; i++ {
		go func() {
			defer wg.Done()
			ageChan <- i
		}()
	}
	wg.Wait()
	close(ageChan)

	ageList := []int{}
	for age := range ageChan {
		ageList = append(ageList, age)
	}
	fmt.Println("age list", ageList)
}

// fibonacci
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func concurrentFib(n int) {
	fib := make(chan int, n)

	go fibonacci(n, fib)

	result := make([]int, 0, n)
	for f := range fib {
		result = append(result, f)
	}
	fmt.Println(result)
}

func channelRunner() {
	simulateEmailSend()
	channels()
	concurrentFib(10)
}
