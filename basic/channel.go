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

func channelRunner() {
  simulateEmailSend()

}
