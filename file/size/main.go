package main

import (
	"bytes"
	"os"
	"strconv"
)

type S struct {
	Id    string
	Value int64
}

func main() {
	// buf := new(bytes.Buffer)
	// for i := 10000000; i < 20000000; i++ {
	// 	// s := S{Id: strconv.Itoa(i), Value: 4200000000 + int64(i)}
	// 	if err := binary.Write(buf, binary.LittleEndian, 4200000000+int64(i)); err != nil {
	// 		panic(err)
	// 	}

	// 	if err := binary.Write(buf, binary.LittleEndian, 4200000000+int64(i)); err != nil {
	// 		panic(err)
	// 	}

	// }
	var buf bytes.Buffer
	for i := 0; i < 12000000; i++ {
		buf.WriteString(strconv.Itoa(i+10000000) + "\n")
	}

	file, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
