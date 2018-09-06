package main

import "fmt"

type I interface {
	M()
	K()
}

type T struct {
	S string
}
type T2 struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S, "M")
}

//func (t T) K() {
//	fmt.Println(t.S, "K")
//}

func (t T) K() {
	fmt.Println(t.S, "K")
}

func main() {
	var i I = T{"hello"}
	i.M()
	i.K()
}
