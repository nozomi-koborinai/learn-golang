package main

import (
	"fmt"
)

func main() {
	var s string
	err := setString(&s, "foo")
	fmt.Println(s, err == nil) // foo true

	var n int
	err = setString(&n, "foo")
	fmt.Println(n, err == nil) // 0 false

}

func setString(val interface{}, str string) error {
	// 型アサーションを使用して、valが参照string型かどうか判定する
	pointer, result := val.(*string)
	if !result {
		return fmt.Errorf("assignString: expected *string, but got %T", val)
	}

	// ポインタがnilかどうか判定する
	if pointer == nil {
		return fmt.Errorf("assignString: expected *string, but got %T", val)
	}

	*pointer = str
	return nil
}
