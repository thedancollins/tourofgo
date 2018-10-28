package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Here")
	fmt.Println(<-ch)
	ch <- 5
	fmt.Println("There")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
