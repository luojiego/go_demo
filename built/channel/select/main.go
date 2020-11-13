package main

func unBufferChannel() {
	println("unBufferChannel")
	var ch = make(chan int)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			println(v)
		}
	}
}

func bufferChannel() {
	println("bufferChannel")
	var ch = make(chan int, 10)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			println(v)
		}
	}
}

func main() {
	unBufferChannel()
	bufferChannel()
}
