package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
)

func main() {
	file, err := excelize.OpenFile("销售线索20200916.xlsx")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rows := file.GetRows("Sheet1")
	for k, row := range rows {
		// fmt.Println(k, row)
		if k > 1 {
			num := row[2]
			result := sha256.Sum256([]byte(num))
			fmt.Printf("%s %x\n", num, result)
			f.WriteString(fmt.Sprintf("%x\n", result))
		}
	}
}
