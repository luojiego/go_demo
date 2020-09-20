package main

import (
	"bytes"
	"encoding/binary"
	"os"
)

type S struct {
	Id    string
	Value int64
}

func main() {
	buf := new(bytes.Buffer)
	for i := 10000000; i < 20000000; i++ {
		// s := S{Id: strconv.Itoa(i), Value: 4200000000 + int64(i)}
		if err := binary.Write(buf, binary.LittleEndian, 4200000000+int64(i)); err != nil {
			panic(err)
		}

		if err := binary.Write(buf, binary.LittleEndian, 4200000000+int64(i)); err != nil {
			panic(err)
		}

	}

	file, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
