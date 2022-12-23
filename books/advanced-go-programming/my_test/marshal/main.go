package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	t := struct {
		time.Time
		N int
	}{time.Now(), 5}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}
