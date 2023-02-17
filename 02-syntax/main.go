package main

import "fmt"

func main() {
	/// varの省略
	// 関数内のみでしかできない
	n := 100
	// まとめる
	var (
		m = 200
		l = 300
		//k = 400	//変数定義したのに使ってない場合はコンパイルエラー
	)
	fmt.Println(n, m, l)

	/// 右辺の省略
	// 省略された定数定義の右辺は、最後の省略されていない定数の右辺と同じになる
	const (
		a = 1 + 2
		b
		c
	)
	fmt.Println(a, b, c)

	/// iota
	const (
		StatusOK = iota
		StatusNG
	)
	fmt.Println(StatusOK, StatusNG)

	/// iotaを利用したフラグ
	// iota分だけ左シフトさせるとiota番目のビットだけ1になる
	// ビット演算するのに便利
	const (
		FlagA = 1 << iota // 1
		FlagB             // 2 （FlagB = 1 << iotaと同じ意味）
		FlagC             // 4
		FlagD             // 8
	)
	fmt.Println(FlagA, FlagB, FlagC, FlagD)

	/// if構文
	// ()がいらない
	if a == 0 {
	}
	// できない①（if行末尾に中括弧がないとコンパイルエラー）
	//if a == 0
	//{
	//}
	// できない②（Dartみたいに１行で書いちゃうとコンパイルエラー）
	//if (a == 0) fmt.Println(a)
	// 代入文を書く
	if a := 3; a > 0 {
		fmt.Println(a)
	} else {
		fmt.Println(2 * a)
	}

	/// switch構文
	// javaのようなケースごとのbreakは不要
	switch a {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}
	// caseに式が使える
	switch {
	case a == 1:
		fmt.Println("a is 1")
	}

	// for構文
	// Goにおいて繰り返しはforしかない
	// do, whileに対応する機能はない
	//  → コードが難解になりやすいため
	// ①初期値;継続条件;更新文
	for i := 0; i <= 100; i = i + 1 {
	}
	// ②継続条件のみ
	//for i <= 100 {
	//}
	// ③無限ループ
	//for {
	//}
	// ④rangeを使った繰り返し
	for _, v := range []int{10, 20, 30} {
		//fmt.Println(i)
		fmt.Println(v)
	}
}
