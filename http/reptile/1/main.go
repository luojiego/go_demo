package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func download(url string)  {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36 Edg/83.0.478.61")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(all))

}

func main() {
	file, err := os.Open("./daily.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		contains := bytes.Contains(line, []byte("magnet:"))
		if contains {
			fmt.Println(string(line))
			split := bytes.Split(line, []byte("href="))
			fmt.Println(len(split))

			s1 := bytes.Split(split[1], []byte("\""))
			fmt.Println(len(s1))
			url := string(s1[1])
			download(url)
			return
		}
	}

	// content, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	/*split := bytes.Split(content, []byte("href=\""))
	fmt.Println("-------------------", len(split))
	for _, v := range split {
		t := bytes.Split(v, []byte("\""))
		fmt.Println(string(t[0]))
		if bytes.HasPrefix(v, []byte("magnet:")) {
			t := bytes.Split(v, []byte("\""))
			fmt.Println(string(t[0]))
			return
		}
	}*/
}
