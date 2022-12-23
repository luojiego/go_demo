package main

var (
	done = make(chan bool)
	msg = ""
)

func aGoroutine() {
	msg = "hello, world"
	close(done)
}

func main() {
	go aGoroutine()
	<- done
	println(msg)
}
