package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Halo dunia!")
	b := make([]byte, 6)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Print(string(b[:n]), "\n")
	}
}
