package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	randSeed = time.Now().UnixNano()
	rnd      *rand.Rand
)

func main() {
	// 乱数の初期化
	rnd = rand.New(rand.NewSource(randSeed))

	// HTTPサーバの作成
	http.HandleFunc("/", omikujiHandler)
	http.ListenAndServe(":8080", nil)
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	result := drawOmikuji()
	fmt.Fprintf(w, "今日の運勢は%sです。\n", result)
}

func drawOmikuji() string {
	// おみくじの結果をランダムに決定
	result := ""
	switch rnd.Intn(6) {
	case 0:
		result = "大吉"
	case 1, 2:
		result = "中吉"
	case 3, 4:
		result = "小吉"
	default:
		result = "吉"
	}
	return result
}
