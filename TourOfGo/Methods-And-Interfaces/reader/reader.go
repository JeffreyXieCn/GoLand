package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
	"strings"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(arr []byte) (int, error) {
	numOfBytes := len(arr)
	for i := 0; i < numOfBytes; i++ {
		arr[i] = 'A'
	}

	return numOfBytes, nil
}

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	fmt.Println("\nImplement a Reader type that emits an infinite stream of the ASCII character 'A'")
	reader.Validate(MyReader{})
}
