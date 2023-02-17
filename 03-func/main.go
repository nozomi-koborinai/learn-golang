package main

import "fmt"

func main() {
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	num := 4.5
	if avg > int(num) {
		println("good")
	}

	var a, b, c bool
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}

	/// 配列の初期化
	// ゼロ値で初期化
	var ns1 [5]int
	// 配列リテラルで初期化
	var ns2 = [5]int{10, 20, 30, 40, 50}
	// 要素数を値から推論
	ns3 := [...]int{10, 20, 30, 40, 50}
	// 5番目が50、10番目が100で他が0の要素数11の配列
	ns4 := [...]int{5: 50, 10: 100}
	fmt.Println(ns1, ns2, ns3, ns4)

	p := struct {
		name string
		age  int
	}{name: "Gopher", age: 10}
	p.age += 1
	fmt.Println(p.name, p.age)

	/// 配列の操作
	ns := []int{10, 20, 30, 40, 50}
	// 要素にアクセス
	println(ns[3])
	// 長さ
	println(len(ns))
	// 容量
	println(cap(ns))
	// 要素の追加
	// 容量が足りない場合は背後の配列が再確保される
	ns = append(ns, 60, 70)
	println(len(ns), cap(ns)) // 長さと容量

	/// マップの操作
	m := map[string]int{"x": 10, "y": 20}
	// キーを指定してアクセス
	println(m["x"])
	// キーを指定して入力
	m["z"] = 30
	// 存在を確認する
	n, ok := m["z"]
	println(n, ok)
	// キーを指定して削除する
	delete(m, "z")
	// 削除されていることを確認
	n, ok = m["z"] // ゼロ値とfalseを返す
	fmt.Println(n, ok)

	x, y := swap(10, 20)
	fmt.Println(x, y)

	/// 無名関数（クロージャ）
	msg := "Hello, 世界"
	func() {
		println(msg)
	}()

	/// 関数型
	fs := make([]func() string, 2)
	fs[0] = func() string { return "hoge" }
	fs[1] = func() string { return "fuga" }
	for _, f := range fs {
		fmt.Println(f())
	}

	/// 値のコピー
	// コピーなのでコピー後に変更を加えてもコピー前には影響しない
	// その逆も然り
	h1 := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}
	h2 := h1 // コピー
	h2.age = 20
	println(h1.age, h1.name)
	println(h2.age, h2.name)

	/// ポインタ
	var x2 int
	f(&x2)
	println(x2)

	// メソッド値・メソッド式
	var hex Hex = 100
	f := hex.String
	g := Hex.String
	fmt.Println(f())
	fmt.Println(g(2))
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

// / ポインタ
func f(xp *int) {
	*xp = 100
}
func swap(x, y int) (int, int) {
	return y, x
}

type Person struct {
	Name string
	Age  int
}
