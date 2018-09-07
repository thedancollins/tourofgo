package main

import "fmt"

func main() {
	var i interface{} //= "hello"

	if i != nil {
		s := i.(string)
		fmt.Println(s)
	}
	i = "Hello"

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	if ok {
		f = i.(float64) // panic
		fmt.Println(f)
	}
}
