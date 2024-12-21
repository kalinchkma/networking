package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutexes(t *testing.T) {
	type testCase struct {
		email string
		count int
	}

	var tests = []testCase{
		{"naruto@konoha.com", 23},
		{"natsu@fairy.com", 30},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.Mutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()

		if output := sc.val(test.email); output != test.count {
			failCount++
			t.Errorf(`
				---------------------------------
				Test Failed:
				email: %v
				count: %v
				expected count: %v
				actual count:   %v
				`, test.email, test.count, test.count, output)
		} else {
			passCount++
			fmt.Printf(`
				---------------------------------
				Test Passed:
				email: %v
				count: %v
				expected count: %v
				actual count:   %v
				`, test.email, test.count, test.count, output)
		}

	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
