package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) print() {
	if p == nil {
		fmt.Println("p is nil")
	}
}

func main() {
	var p *Person
	p.print()

	m := make(map[int]Person)
	m[1] = Person{
		Name: "Roger",
		Age:  30,
	}
	// m[1].Name = "22" // cannot assign to struct field m[1].Name in map

	m1 := make(map[int]*Person)
	m1[1] = &Person{
		Name: "罗杰",
		Age:  30,
	}

	for _, v := range m1 {
		v.Age = 15
		test(v)
	}
}

func test(p *Person) {
	fmt.Println(p)
}
