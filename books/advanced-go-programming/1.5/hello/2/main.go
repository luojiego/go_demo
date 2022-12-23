package main

func main() {
	done := make(chan int)

	go func() {
		println("hello world")
		done <- 1
	}()
	<- done
}
