package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch, quit chan int) {  //
	fmt.Println(t);
	for {
		select {
		case ch <- t.Value:
			if t.Left != nil {
				//fmt.Println(t.Left);
				go Walk(t.Left, ch, quit)
			}
			if t.Right != nil { //
				//fmt.Println(t.Right);
				go Walk(t.Right, ch, quit)
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool

func main() {
  ch := make(chan int)
  quit := make(chan int)
  go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
  Walk(tree.New(1), ch, quit)
}
