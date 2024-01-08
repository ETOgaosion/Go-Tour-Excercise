package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	lastRes := 0
	res := 1
	index := 0
	return func() int {
		if index < 2 {
			index++
			return index - 1
		} else {
			tmp := res
			res += lastRes
			lastRes = tmp
			return res
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
