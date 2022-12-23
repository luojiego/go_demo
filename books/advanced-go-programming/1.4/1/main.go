package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	file, err := os.Create("test.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello world")
	fmt.Fprintln(&UpperWriter{file}, "hello world")
}
