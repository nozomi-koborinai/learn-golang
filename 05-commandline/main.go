package _5_commandline

import "fmt"

func main() {
	/// deferは関数終了後に評価される
	msg := "!!!"
	defer fmt.Println(msg)
	msg = "world"
	defer fmt.Println(msg)
	fmt.Println("hello")
}
