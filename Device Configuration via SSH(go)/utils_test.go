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

func TestReadIPFile(t *testing.T) {
	errorFileName := "./ajkkd.txt"
	validFileName := "./ip_test.txt"
	if _, err := ReadIPFile(errorFileName); err == nil {
		t.Fatalf("Error Reading IP File")
	}
	if ipList, err := ReadIPFile(validFileName); err != nil {
		t.Fatalf("Error Reading IP File")
	} else {
		expectedIpList := []string{"192.168.1.11", "192.168.1.12", "192.168.1.13"}
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
