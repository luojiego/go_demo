package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u User) PrintAge() {
	fmt.Println(u.Age)
}

func test(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println("v Name: ", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("the argument can't equal struct")
		return
	}

	value := reflect.ValueOf(v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := value.Field(i).Interface()
		fmt.Printf("%8s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		f := t.Method(i)
		fmt.Printf("%16s: %v\n", f.Name, f.Type)
	}
}

func main() {
	u := User{"LuoJie", 30, "Xi'an"}
	test(u)
}
