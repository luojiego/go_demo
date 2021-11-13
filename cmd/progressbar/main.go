package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func main() {
	p1()
	p2()
	p3()
}

func p1() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}

func p2() {
	req, _ := http.NewRequest("GET", "https://dl.google.com/go/go1.14.2.src.tar.gz", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile("go1.14.2.src.tar.gz", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}

func p3() {
	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan][1/3][reset] Writing moshable file..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	for i := 0; i < 1000; i++ {
		bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}
}
