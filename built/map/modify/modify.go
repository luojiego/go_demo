package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	m := make(map[string]*Person, 0)

	m["一"] = &Person{
		Name: "1",
	}
	fmt.Println(m)

	v, ok := m["一"]
	if ok {
		// v = &Person{
		// 	Name: "1111",
		// }
		v.Name = "1111"
	}
	_ = v

	/*if _, ok := m["一"]; ok {
		m["一"] = &Person{
			Name: "1111",
		}
	}*/

	for k, v := range m {
		fmt.Println(k, v.Name)
	}
}
