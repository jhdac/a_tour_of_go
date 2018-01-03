package main

import "fmt"

type IPAddr [4]byte

func (s IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", s[0], s[1], s[2], s[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
