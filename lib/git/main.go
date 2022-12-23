package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	path := os.Getenv("SynthesisAndCoc")
	fmt.Println(path)

	// 指定子仓库路径
	path += string(os.PathSeparator) + "pb" + string(os.PathSeparator) + "proto" + string(os.PathSeparator)

	res, err := git.PlainOpen(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	filename := "data/data.proto"
	// 指定文件
	commiter, err := res.Log(&git.LogOptions{
		FileName: &filename,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	first, err := commiter.Next()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(first.Hash)
	fmt.Println(first.Author.When.Format("2006-01-02 15:04:05"))
	fmt.Println(first.Message)
}
