package main

func main() {
	done := make(chan int)

	go func() {
		println("hello world")
		<-done
	}()
	done <- 1
}
