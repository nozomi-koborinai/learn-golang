package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	words := []string{"apple", "banana", "cherry", "durian", "elderberry", "fig", "grape"}

	rand.Seed(time.Now().UnixNano())
	score := 0
	start := time.Now()

	// 制限時間を設定する
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 1; i <= 10; i++ {
		word := words[rand.Intn(len(words))]
		fmt.Printf("Type the word: %s\n", word)

		answerCh := make(chan string)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			answerCh <- text
		}()

		// 制限時間と入力を待つ
		select {
		case <-ctx.Done():
			fmt.Println("Time's up!")
		case answer := <-answerCh:
			if answer == word {
				score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect.")
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Score: %d / 10\nTime: %s\n", score, elapsed)
}
