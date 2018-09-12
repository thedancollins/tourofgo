package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type ArchMage struct {
	person Person
	level  int
}

func (p ArchMage) String() string {
	return fmt.Sprintf("%v (Level %v)", p.person, p.level)
}

func main() {
	a := Person{"Arthur Dent", 42}
	y := Person{"LeJames Bron", 30}
	z := ArchMage{Person{"Zaphod Beeblebrox", 9001}, 9002}
	fmt.Println(a, y, z)
}
