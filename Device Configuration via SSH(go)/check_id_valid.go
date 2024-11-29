package main

import (
	"net"
	"strconv"
	"strings"
)

// Validate IP Address
//
// # False
//
// If not valid
func ValidateIPAddress(ip string) bool {
	address := net.ParseIP(ip)
	if address != nil {
		addressList := strings.Split(ip, ".")
		octectList := make([]uint8, 4)
		for i, octect := range addressList {
			if n, err := strconv.Atoi(octect); err != nil {
				return false
			} else {
				octectList[i] = uint8(n)
			}
		}
		if octectList[0] <= 223 && octectList[0] != 127 && (octectList[0] != 169) || (octectList[1] != 254) && octectList[1] < 255 && octectList[2] < 255 && octectList[3] < 255 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
