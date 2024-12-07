package main

// import (
// 	"fmt"
// 	"time"
// )

// type request struct {
// 	path string
// }

// func main() {
// 	reqs := make(chan request, 100)
// 	go handleRequests(reqs)
// 	for i := 0; i < 4; i++ {
// 		reqs <- request{path: fmt.Sprintf("/path/%d", i)}
// 		time.Sleep(500 * time.Microsecond)
// 	}
// 	// time.Sleep(5 * time.Second)
// 	fmt.Println("5 seconds passed, killing server")
// }

// func handleRequests(reqs <-chan request) {
// 	for req := range reqs {
// 		go handleRequest(req)
// 	}
// }

// func handleRequest(req request) {
// 	fmt.Println("Handling request for", req.path)
// 	time.Sleep(time.Second * 2)
// 	fmt.Println("Done with request for", req.path)
// }
