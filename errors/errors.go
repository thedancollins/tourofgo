package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When  time.Time
	What  string
	other string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s, %s",
		e.When, e.What,
		e.other)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
		"Yes it did",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
