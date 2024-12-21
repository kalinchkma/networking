package main

import (
	"fmt"
	"time"
)

func RunProcess() {
	messages := []string{
		"Hey buddy whats up",
		"Long time no see",
		"What's up",
		"I am fine what about you",
	}
	startTime := time.Now()
	fmt.Println(processMessages(messages))
	endTime := time.Since(startTime)
	fmt.Printf("Time taken %vs\n", endTime)
}

func processMessages(messages []string) []string {
	msg := make(chan string, len(messages))
	fmt.Println("process started")

	for _, m := range messages {
		go func() {
			msg <- process(m)
		}()
	}
	processMsg := []string{}
	for i := 0; i < len(messages); i++ {
		processMsg = append(processMsg, <-msg)
	}
	return processMsg
}

func process(message string) string {
	time.Sleep(5 * time.Second)
	return message + "-processed"
}
