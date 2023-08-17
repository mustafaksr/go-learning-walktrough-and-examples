package main

import "reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A' // Fill the byte slice with ASCII 'A' character
	}
	return len(b), nil // Return the number of bytes read and nil error
}

func main() {
	reader.Validate(MyReader{})
}
