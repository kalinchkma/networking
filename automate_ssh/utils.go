package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Validate IP Address
//
// # False
//
// If not valid
func ValidateIPAddress(ip string) bool {
	// Convert string to valid ip address
	address := net.ParseIP(ip)
	// Check if ip address is valid or not
	if address != nil {
		// Split ip address into string octects list
		octetList := strings.Split(ip, ".")
		// Unint8 octets holders
		octets := make([]uint8, 4)
		// Loop throught the octets to convert string to unint8
		for i, octect := range octetList {
			// Onvert octets into number
			if n, err := strconv.Atoi(octect); err != nil {
				// If octet cannot convert into number return invalid (false)
				return false
			} else {
				// Push converted octet into octets holder
				octets[i] = uint8(n)
			}
		}
		// Additional check for ip address
		if octets[0] <= 223 && octets[0] != 127 && (octets[0] != 169) || (octets[1] != 254) && octets[1] < 255 && octets[2] < 255 && octets[3] < 255 {
			// If ip address octet dose meet requirement return valid (true)
			return true
		} else {
			// If ip address octet dose not meet requirement return invalid (false)
			return false
		}
	} else {
		// If ip address is not valid return invalid (false)
		return false
	}
}

// Content type enum for reading file
type ContentType string

const (
	ContentTypeString ContentType = "string"
	ContentTypeList   ContentType = "list"
)

// Read IP File
//
// # Error
//
// If file not exist and cannot read file
func ReadFile(ipFileName string, contentType ContentType) (interface{}, error) {
	// Check if file is exist
	if file, err := os.OpenFile(ipFileName, os.O_RDONLY, os.ModeDevice); err != nil {
		// Provided file dose not exist
		return nil, err
	} else {
		// Close the file when done reading
		defer file.Close()
		// Read the file content
		if fileContent, err := io.ReadAll(file); err != nil {
			// Error reading file
			return nil, err
		} else {
			// Check which type to return
			switch contentType {
			case ContentTypeList:
				// Return slice of line string
				return strings.Split(string(fileContent), "\n"), nil
			case ContentTypeString:
				// Retrun content of string
				return string(fileContent), nil
			default:
				// Return slice of byte
				return fileContent, nil
			}
		}
	}
}

// Ping host machine by IP address
//
// # False
//
// If host not reachable by ping request
func PingHostByIP(ip string, tt int) bool {
	// Ping the host
	cmd := exec.Command("ping", ip, "-c", "2", "-t", fmt.Sprintf("%v", tt))

	// Pin command output to standard output
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		// If host cannot reach by a ping, return error
		return false
	} else {
		// If host reach able by a ping, return true
		return true
	}
}

// Credentials type
type Credential struct {
	username string
	password string
}

// Parse ssh credentials from file
//
// # List of credentials, error
//
// Return array of ssh credentials or error
func GetSSHCredentials(filePath string) ([]Credential, error) {
	// Read Credentials data from file
	if data, err := ReadFile(filePath, "list"); err != nil {
		// Return file reading error
		return nil, err
	} else {
		// Credentials list holder
		var credentials []Credential
		// Convert data into list of data
		dataList := data.([]string)
		// Loop through the data list
		for _, c := range dataList {
			// Split the data line by `,`
			newCred := strings.Split(c, ",")
			// Check line contains username and password
			if len(newCred) != 2 {
				// Return error if line dose not contain username and password
				return nil, fmt.Errorf("invalid file format %v", c)
			}
			// Serialized the line credentials and push it to credentials holder
			credentials = append(credentials, Credential{
				username: newCred[0],
				password: newCred[1],
			})
		}
		// Return credential list
		return credentials, nil
	}
}
