package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age int
}

/*func (u *User) String() string {
	result, _ := json.Marshal(u)
	//fmt.Println(string(result))
	return string(result)
}*/

func (u User) String() string {
	result, _ := json.Marshal(u)
	//fmt.Println(string(result))
	return string(result)
}

func main() {
	u := User{
		Name: "luojie",
		Age: 30,
	}

	fmt.Printf("%s\n", &u)
	fmt.Println(u)
}
