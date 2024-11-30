package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

func main() {
	data, err := ReadFile("ip.txt", "list")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	ipList := data.([]string)
	credentials, err := GetSSHCredentials("credentials.csv")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	cmdData, err := ReadFile("command.txt", "list")
	cmds := strings.Join(cmdData.([]string), " && ")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if len(ipList) != len(credentials) {
		log.Fatalf("Invalid ip and credentials")
	}

	var wg sync.WaitGroup

	wg.Add(len(ipList))
	for i := 0; i < len(ipList); i++ {
		go func() {
			defer wg.Done()
			SSHConnection(ipList[i], credentials[i], cmds)
		}()
	}

	wg.Wait()
	fmt.Println("Done proccessing,.....")
}
