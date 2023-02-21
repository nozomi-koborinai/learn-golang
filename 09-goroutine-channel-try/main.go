package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	// チャネルを作る
	out := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// チャネルに、読み込んだ文字列を送る
			out <- s.Text()
		}
		// チャネルを閉じる
		close(out)
	}()
	// チャネルを返す
	return out
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
