package main

import (
	"fmt"
    "strconv"
)

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ipAddr IPAddr) String() string {
	var res string
	for idx, num := range ipAddr {
		res += strconv.Itoa(int(num))
		if idx != 3 {
			res += "."
		}
	}
	return res
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
