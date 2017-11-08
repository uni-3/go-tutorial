package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	var n int
	var err error
	for n, err = 0, nil; n < len(b); n++ {
		b[n] = 'A'
	}
	return n, err
}

func main() {
	reader.Validate(MyReader{})
}
