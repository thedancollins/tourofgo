package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Printf("The map: %v.  The value: %d.\n", m, m["Answer"])

	m["Answer"] = 48
	fmt.Printf("The map: %v.  The value: %d.\n", m, m["Answer"])

	delete(m, "Answer")
	fmt.Printf("The map: %v.  The value: %d.\n", m, m["Answer"])

	v, ok := m["Answer"]
	fmt.Printf("The map: %v.  The value: %d. Present? %t\n", m, v, ok)
}
