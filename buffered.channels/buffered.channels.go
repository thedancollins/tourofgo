package main

import (
	"fmt"
	"time"
)

func say(s chan int) {
	//for i := 0; i < 5; i++ {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(<-s)
	//}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Here1")
	go say(ch)
	fmt.Println("Here2")
	ch <- 5
	fmt.Println("Here3")
	fmt.Println(<-ch)
	fmt.Println("There")
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
