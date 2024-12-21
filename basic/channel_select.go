package main

import (
	"fmt"
	"math/rand"
	"time"
)

func logger() {
	sms := []string{
		"First message",
		"Second Message",
		"Thirth Message",
	}
	emails := []string{
		"naruto@konoha.com",
		"uchiha@konoha.com",
		"obito@pain.com",
	}

	fmt.Println("Starting logging....")
	chSms, chEmails := sendToLogger(sms, emails)
	for {
		select {
		case s, ok := <-chSms:
			if !ok {
				chSms = nil
			} else {
				fmt.Println("SMS:", s)
			}
		case e, ok := <-chEmails:
			if !ok {
				chEmails = nil
			} else {
				fmt.Println("EMAIL:", e)
			}
		}

		if chSms == nil && chEmails == nil {
			break
		}
	}
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			// done channel is used to ensure that both sending operations completed
			// before moving next iteration
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()

			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()

			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return
}
