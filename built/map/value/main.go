package main

type Person struct {
	Name string
	Age  int
}

func main() {
	m := make(map[int]Person)
	m[1] = Person{
		Name: "Roger",
		Age:  30,
	}
	m[1].Name = "22" // cannot assign to struct field m[1].Name in map

}
