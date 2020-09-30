package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
)

var (
	filename    = ""
	sheetName   = ""
	srcRow      = ""
	dstRow      = ""
	upperResult = false
)

func init() {
	flag.StringVar(&filename, "filename", "test.xlsx", "目标文件名称")
	flag.StringVar(&sheetName, "sheet", "Sheet1", "excel 标签页[一定要确保正确]")
	flag.StringVar(&srcRow, "src", "F", "目标列[左右的空格会被移除]")
	flag.StringVar(&dstRow, "dst", "G", "md5 结果写入列")
	flag.BoolVar(&upperResult, "upperResult", false, "md5 结果转成大写")
}

func main() {
	flag.Parse()
	file, err := excelize.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(srcRow, []byte(srcRow)[0]-'A')

	// return
	rows := file.GetRows(sheetName)

	for k, row := range rows {
		if k > 0 {
			s := strings.TrimSpace(row[[]byte(srcRow)[0]-'A'])
			d := ""
			if upperResult {
				d = fmt.Sprintf("%X", md5.Sum([]byte(s)))
			} else {
				d = fmt.Sprintf("%x", md5.Sum([]byte(s)))
			}
			file.SetCellValue(sheetName, dstRow+strconv.Itoa(k+1), d)
			fmt.Println(s, d, dstRow+strconv.Itoa(k))
		}
	}
	file.Save()
}
