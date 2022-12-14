package main

import "os"

func main() {
	ParalleWrite([]byte{0x31, 0x23, 0x56})
}

// 第一次见 chan error，看来我还是太年轻了
func ParalleWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}

	f2, err := os.Create("file2")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}
