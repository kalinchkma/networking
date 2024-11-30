package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

// Connect to ssh and automate your needs
//
// # Run the automation
//
// Its return nothing
func SSHConnection(ip string, credentials Credential, command string) {
	// Configure ssh
	sshConfig := &ssh.ClientConfig{
		User: credentials.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(credentials.password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Accept any host key
	}

	// Establish the connection
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:22", ip), sshConfig)
	if err != nil {
		log.Fatalf("Faild to connect to SSH server: %v", err)
	}
	defer client.Close()
	fmt.Printf("SSH connection established to %v\n", credentials.username)

	// Create a ssh session
	session, err := client.NewSession()

	if err != nil {
		log.Fatalf("Error creating new session %v", err)
	}

	defer session.Close()

	// execute remote command
	if res, err := session.CombinedOutput(command); err != nil {
		log.Fatalf("Failed to execute command: %v: %v", ip, err)
	} else {
		fmt.Println(credentials.username, res)
	}

}
