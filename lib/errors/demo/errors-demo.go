package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
)

func fun1() error {
	return fun2()
}

func fun2() error {
	_, err := os.Open("123")
	if err != nil {
		return err
	}
	return errors.Wrap(err, "open failed")
}

func main() {
	err := fun1()
	if err != nil {
		log.Println(err)
	}
	// fun1:fun2: open 123: The system cannot find the file specified.
}
