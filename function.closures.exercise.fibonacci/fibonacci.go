package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last := 0
	prev := 0
	now := 0
	hiti := 0
	return func() int {
		hiti++
		if hiti == 2 || hiti == 3 {
			now = 1
		}
		prev = last
		last = now

		if hiti != 2 && hiti != 3 {
			now = prev + last
		}
		return now
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println("value ", f())
	}
}
