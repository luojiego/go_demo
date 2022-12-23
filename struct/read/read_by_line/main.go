package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	filename = "./ActivityTaskRewardData.csv"
)

func main() {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	index := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		index++
		if index > 1 {
			//buf := string(line)
			//fmt.Println(buf)

			result := strings.Split(string(line), "\"")
			if len(result) != 3 {
				break
			}
			fmt.Println(result)
			sect := strings.Split(result[0], ",")
			if len(sect) != 3 {
				break
			}
			fmt.Println(len(sect), ":\t", sect)

			reward := strings.ReplaceAll(result[1], "[", "")
			reward = strings.ReplaceAll(reward, "]", "")
			fmt.Println(reward)

			rewardResult := strings.Split(reward, ",")
			//fmt.Println(len(rewardResult))
			if len(rewardResult) != 2 {
				break

			}
			fmt.Println(rewardResult)
		}
	}
}
