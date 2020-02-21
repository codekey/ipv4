// For simple operations around IP addresses
package ipv4

import (
	"strconv"
	"strings"
)

// Sort the given IPv4 address (slice of String) in the ascending order in-place
func Sort(reqIP []string) {
	for i := 0; i < len(reqIP)-1; i++ {
		for j := i + 1; j < len(reqIP); j++ {
			for pos := 1; pos <= 4; pos++ {
				x := getOctet(reqIP[i], pos)
				y := getOctet(reqIP[j], pos)

				if x > y {
					t := reqIP[i]
					reqIP[i] = reqIP[j]
					reqIP[j] = t
				}

				if x == y {
					continue
				} else {
					break
				}
			}
		}
	}
}

// Returns true if given IP is public or else vice versa
func IsPublic(reqIP string) bool {
	return !checkIfPrivate(reqIP)
}


// Returns true if given IP is private or else vice versa
func IsPrivate(reqIP string) bool {
	return checkIfPrivate(reqIP)
}

// Returns true if given IP is CGN (Carrier-Grade NAT) or else vice versa
func IsCGN(reqIP string) bool {
	oct1 := getOctet(reqIP, 1)
	oct2 := getOctet(reqIP, 2)

	if oct1 == 100 {
		if oct2 >= 64 && oct2 <= 127 {
			return true
		}
	}

	return false
}

// For internal use - For checking if the given IP is private
func checkIfPrivate(reqIP string) bool {
	oct1 := getOctet(reqIP, 1)
	oct2 := getOctet(reqIP, 2)

	if oct1 == 10 {
		return true
	}

	if oct1 == 172 {
		if oct2 >= 16 && oct2 <= 31 {
			return true
		}
	}

	if oct1 == 192 && oct2 == 168 {
		return true
	}

	if oct1 == 100 {
		if oct2 >= 64 && oct2 <= 127 {
			return true
		}
	}

	return false
}

// For internal use - For obtaining desired octet from given IP
func getOctet(ip string, pos int) int {
	x, _ := strconv.Atoi(strings.Split(ip, ".")[pos-1])
	return x
}
