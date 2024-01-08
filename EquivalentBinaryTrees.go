package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int, outside bool) {
	// fmt.Println(t.Value)
	if t.Left != nil {
		Walk(t.Left, ch, false)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch, false)
	}
	if outside {
		close(ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)
	for {
		i1, ok1 := <-ch1
		i2, ok2 := <-ch2
		// fmt.Println(i1, i2, ok1, ok2)
		if ok1 != ok2 {
			return false
		} else if !ok1 {
			return true
		}
		if i1 != i2 {
			return false
		}
	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("Result: ", Same(t1, t2))
}
