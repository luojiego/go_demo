package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Person")
			return new(Person)
		},
	}
}

func main() {
	initPool()

	p := pool.Get().(*Person)
	fmt.Println("first Get Person From pool: ", p)

	p.Name = "Roger"
	fmt.Printf("set p.Name = %s\n", p.Name)

	pool.Put(p)

	fmt.Println("pool has one object: &{Roger}, pool.Get: ", pool.Get().(*Person))
	fmt.Println("pool no more object, pool.Get: ", pool.Get().(*Person))
}
