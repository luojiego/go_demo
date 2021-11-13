package main

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var (
	filename     = ""
	sheetName    = ""
	srcRow       = ""
	md5DstRow    = ""
	sha256DstRow = ""
	md5Upper     = false
	sha256Upper  = false
)

func init() {
	flag.StringVar(&filename, "filename", "test.xlsx", "目标文件名称")
	flag.StringVar(&sheetName, "sheet", "Sheet1", "excel 标签页[一定要确保正确]")
	flag.StringVar(&srcRow, "src", "A", "目标列[左右的空格会被移除]")
	flag.StringVar(&md5DstRow, "md5_dst", "", "md5 结果写入列，不配置则会新增一列写入")
	flag.StringVar(&sha256DstRow, "sha_256_dst", "", "sha256 结果写入列，不配置则会新增一列写入")
	flag.BoolVar(&md5Upper, "md5_upper", false, "md5 结果转成大写，默认为 false")
	flag.BoolVar(&sha256Upper, "sha256_upper", false, "sha256 结果转成大写，默认为 false")
}

func Index2ExcelRow(index int) string {
	var Letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	result := Letters[index%26]
	index = index / 26
	for index > 0 {
		index = index - 1
		result = Letters[index%26] + result
		index = index / 26
	}
	return result
}

func main() {
	flag.Parse()
	file, err := excelize.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	// fmt.Println(srcRow, []byte(srcRow)[0]-'A')

	// return
	rows := file.GetRows(sheetName)
	if len(rows) == 0 {
		fmt.Println("无法为空文档进行加密")
		return
	}

	var srcValue string
	for k, row := range rows {
		// 判断是否要更新
		if k == 0 {
			srcValue = file.GetCellValue(sheetName, srcRow+strconv.Itoa(1))
			writeCellIndex := len(row)
			// fmt.Printf("first row len: %d\n", len(row))
			if md5DstRow == "" {
				md5DstRow = Index2ExcelRow(writeCellIndex)
				writeCellIndex++
				// fmt.Printf("md5DstRow: %s\n", md5DstRow)
				style, _ := file.NewStyle(`{"font":{"family":"Consolas","size":11,"color":"#000000"}}`)
				file.SetCellValue(sheetName, md5DstRow+strconv.Itoa(1), "md5("+srcValue+")")
				file.SetCellStyle(sheetName, md5DstRow+strconv.Itoa(1), md5DstRow+strconv.Itoa(len(rows)), style)
				file.SetColWidth(sheetName, md5DstRow, md5DstRow, 34)
			}

			if sha256DstRow == "" {
				sha256DstRow = Index2ExcelRow(writeCellIndex)
				writeCellIndex++
				// fmt.Printf("sha256DstRow: %s --> %s\n", sha256DstRow, sha256DstRow+strconv.Itoa(len(rows)))
				style, _ := file.NewStyle(`{"font":{"family":"Consolas","size":11,"color":"#000000"}}`)
				file.SetCellValue(sheetName, sha256DstRow+strconv.Itoa(1), "sha256(md5("+srcValue+"))")
				file.SetColWidth(sheetName, sha256DstRow, sha256DstRow, 66)
				file.SetCellStyle(sheetName, sha256DstRow+strconv.Itoa(1), sha256DstRow+strconv.Itoa(len(rows)), style)
			}
		}

		if k > 0 {
			s := strings.TrimSpace(row[[]byte(srcRow)[0]-'A'])
			md5Result, sha256Result := "", ""
			if md5Upper {
				md5Result = fmt.Sprintf("%X", md5.Sum([]byte(s)))
			} else {
				md5Result = fmt.Sprintf("%x", md5.Sum([]byte(s)))
			}
			// fmt.Println(md5Result)

			if sha256Upper {
				sha256Result = fmt.Sprintf("%X", sha256.Sum256([]byte(md5Result)))
			} else {
				sha256Result = fmt.Sprintf("%x", sha256.Sum256([]byte(md5Result)))
			}
			// fmt.Println(sha256Result)

			file.SetCellValue(sheetName, md5DstRow+strconv.Itoa(k+1), md5Result)
			file.SetCellValue(sheetName, sha256DstRow+strconv.Itoa(k+1), sha256Result)

			// fmt.Println(s, md5Result, md5DstRow+strconv.Itoa(k+1))
			// fmt.Println(s, sha256Result, sha256DstRow+strconv.Itoa(k+1))
		}
	}
	file.Save()
}
