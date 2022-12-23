package h2

import (
	"build/init/h1"
)

func init() {
	println("I'm h2")
}

func Hello() {
	h1.Hello()
	println("hello h2")
}
