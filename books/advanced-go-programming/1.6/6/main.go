package main

import "fmt"

func GenerateNatural() chan int  {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <-i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	fmt.Println("prime: ", prime)
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i % prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime) //1 2
		ch = PrimeFilter(ch, prime)//3 2
	}
}
