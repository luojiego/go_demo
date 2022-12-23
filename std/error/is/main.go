package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
	err1 := fmt.Errorf("wrap err1: %w", ErrSentinel)
	err2 := fmt.Errorf("warp err2: %w", err1)
	if errors.Is(err2, ErrSentinel) {
		fmt.Println("err2 is ErrSentinel")
		return
	}
	fmt.Println("err2 is not ErrSentinel")
}
