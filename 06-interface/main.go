package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type Func func() string

func (f Func) String() string { return f() }

type Hoge struct{ N int }
type Fuga struct{ Hoge }

func main() {
	// not empty interface
	type Stringer interface {
		String() string
	}
	var s Stringer = Hex(100)
	fmt.Println(s.String())

	// empty interface
	// メソッドセットが空なインターフェース
	// つまりどの型の値も実装していることになる
	// JavaのObject型のような使い方ができる
	var v interface{}
	v = 100
	v = "foo"
	fmt.Println(v)

	// 関数にメソッドを持たせる
	var f fmt.Stringer = Func(func() string {
		return "hi"
	})
	fmt.Println(f)

	// インターフェースの実装チェック
	// コンパイル時に実装しているかチェックする
	// インタフェース型の変数に代入してみる
	var _ fmt.Stringer = Func(nil)

	// 型アサーション
	v = 100
	n, ok := v.(int)
	fmt.Println(n, ok)
	l, ok := v.(string)
	fmt.Println(l, ok)

	// 型スイッチ
	var i interface{}
	i = 100
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "hoge")
	default:
		fmt.Println("default")
	}

	// 埋め込みとフィールド
	fuga := Fuga{Hoge{N: 100}} // Fuga{Hoge:Hoge{N:100}} でもOK
	// Hoge型のフィールドにアクセスできる
	fmt.Println(fuga.N)
	// 型名を指定してアクセスできる
	fmt.Println(fuga.Hoge.N)
}
