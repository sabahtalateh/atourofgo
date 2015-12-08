package main

import "fmt"

type IPAddr [4]byte

func (IP IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", IP[0], IP[1], IP[2], IP[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

