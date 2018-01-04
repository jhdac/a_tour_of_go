package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	"reflect"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	var ch = make(chan int)
	a1 := make([]int, 0)
	go Walk(t1, ch)

	for i := 0; i < 10; i++ {
		a1 = append(a1, <-ch)
	}

	a2 := make([]int, 0)
	go Walk(t2, ch)
	for i := 0; i < 10; i++ {
		a2 = append(a2, <-ch)
	}

	return reflect.DeepEqual(a1, a2)
}

func main() {
	var same bool
	same = Same(tree.New(1), tree.New(1))
	fmt.Printf("The trees are the same: %t\n", same)
	same = Same(tree.New(1), tree.New(2))
	fmt.Printf("The trees are the same: %t\n", same)
}
