package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData("https://api.boot.dev/v1/courses_rest_api/learn-http/issues")
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	data, err := prettify(issues)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	log.Printf("%v", data)

	// read user data
	users, err := getUserData("https://api.boot.dev/v1/courses_rest_api/learn-http/users")
	if err != nil {
		log.Fatalf("error getting user data: %s", err)
	}
	data, err = prettify(users)
	if err != nil {
		log.Fatalf("error decorating users: %s", err)
	}
	fmt.Println(data)
}
