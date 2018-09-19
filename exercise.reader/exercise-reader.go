package main

import "golang.org/x/tour/reader"
import "fmt"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (n int, err error) {

	var s string = ""

	for i := 0; i < len(b); i++ {
		s += "A"
	}
	fmt.Printf(s)

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
