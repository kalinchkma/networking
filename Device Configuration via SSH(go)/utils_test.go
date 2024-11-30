package main

import (
	"fmt"
	"testing"
)

func TestValidateIPAddress(t *testing.T) {
	ip := "255.250.0.0"
	ip2 := "1.1.1.1"
	invalidIP := "aw:2d:qw:33"
	if !ValidateIPAddress(ip) {
		t.Fatalf("Invalid ip address")
	}
	if !ValidateIPAddress(ip2) {
		t.Fatalf("Invalid ip address")
	}

	if ValidateIPAddress(invalidIP) {
		t.Fatalf("Invalid ip address")
	}
}

func TestReadFile(t *testing.T) {
	errorFileName := "./ajkkd.txt"
	validFileName := "./ip_test.txt"
	if _, err := ReadFile(errorFileName, ""); err == nil {
		t.Fatalf("Error Reading File")
	}
	if data, err := ReadFile(validFileName, "list"); err != nil {
		t.Fatalf("Error Reading File")
	} else {
		expectedIpList := []string{"192.168.1.11", "192.168.1.12", "192.168.1.13"}
		ipList := data.([]string)
		if len(ipList) != len(expectedIpList) {
			fmt.Println("Faild-----")
			fmt.Printf("Expected: \n%#v\n", expectedIpList)
			fmt.Printf("Found: \n%#v\n", ipList)
			t.Failed()
		}
		for i := 0; i < len(ipList); i++ {
			if expectedIpList[i] != ipList[i] {
				fmt.Println("Faild-----")
				fmt.Printf("Expected: \n%#v\n", expectedIpList[i])
				fmt.Printf("Found: \n%#v\n", ipList[i])
				t.Failed()
			}
		}
	}
}

func TestPingHostByIP(t *testing.T) {
	unAvailableIp := "192.168.21.11"
	availdAbleID := "192.168.1.11"
	if PingHostByIP(unAvailableIp, 2) {
		t.Fatalf("Faild to reach to the ip address")
	}
	if !PingHostByIP(availdAbleID, 2) {
		t.Fatalf("Faild to reach to the ip address")
	}
}

func TestGetSSHCredentials(t *testing.T) {
	validFile := "./credentials_test.csv"
	invalidFile := "./credentials_test_2.csv"

	// Case 1: Invalid file parsing
	if _, err := GetSSHCredentials(invalidFile); err == nil {
		fmt.Println("Case 1: Parse Invaild file must return error")
		t.Fatalf("Invalid file parsing..")
	}

	// Case 2: Valid file parsing
	if data, err := GetSSHCredentials(validFile); err != nil {
		fmt.Println("Case 2: Parse valid file must return data")
		t.Fatalf("Invalid file parsing..")
	} else {
		expected := []Credential{
			{
				username: "root",
				password: "password",
			},
			{
				username: "root",
				password: "",
			},
			{
				username: "root",
				password: "password",
			},
		}
		if len(data) != 3 {
			fmt.Println("Expected", expected)
			fmt.Println("Found", data)
			fmt.Println("Invalid File parsing")
			t.Failed()
		}
		for i := 0; i < 3; i++ {
			if data[i].username != expected[i].username || data[i].password != expected[i].password {
				fmt.Println("Expected", expected[i])
				fmt.Println("Found", data[i])
				fmt.Println("Invalid File parsing")
				t.Failed()
			}
		}
	}

}
