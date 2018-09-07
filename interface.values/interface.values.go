package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

//I noticed that this funcction had a reference type (t *T) so I changed it to (t T) and added the "t.S += " " + t.S" line to verify I knew how it would work
func (t T) M() {
	fmt.Println(t.S)
	t.S += " " + t.S
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()
	//Added to confirm test above.
	describe(i)

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
