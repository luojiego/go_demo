package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	file, err := excelize.OpenFile("./1.xlsx")
	if err != nil {
		panic(err)
	}

	rows := file.GetRows("")
	for k, row := range rows {
		fmt.Println(k, row)
	}
}
