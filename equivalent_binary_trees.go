package main

import (
	"golang.org/x/tour/tree"
	"reflect"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	vals1 := make([]int, 0)
	vals2 := make([]int, 0)
	for x := range ch1 {
		vals1 = append(vals1, x)
	}
	for x := range ch2 {
		vals2 = append(vals2, x)
	}

	sort.Ints(vals1)
	sort.Ints(vals2)
	if reflect.DeepEqual(vals1, vals2) {
		return true
	} else {
		return false
	}
}

func main() {
	println(Same(tree.New(1), tree.New(1)))
}
