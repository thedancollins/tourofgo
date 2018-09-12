package main

import "fmt"

type IPAddr [4]byte

func (p IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", p[0], p[1], p[2], p[3])
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
		"goatse":    {1, 2, 3, 4},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
