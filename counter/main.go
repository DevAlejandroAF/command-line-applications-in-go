package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./text.txt")
	wordCount := CountWords(data)
	fmt.Println("Count:", wordCount)
}

func CountWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	wordCount := 0
	for _, x := range data {
		if x == ' ' {
			wordCount++
		}
	}
	wordCount++
	return wordCount
}
