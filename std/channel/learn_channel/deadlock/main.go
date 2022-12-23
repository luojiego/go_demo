package main

func main() {
    // chan can't use in one goroutine
    ch := make(chan int)
    ch <- 1
    <- ch
}
