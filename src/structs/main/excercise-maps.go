package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		_, contains := wordMap[word]
		if contains {
			wordMap[word]++
			continue
		}
		wordMap[word] = 1
	}

	return wordMap
}

func main() {
	fmt.Println(WordCount("Hello Hello My name is Alex! YES! Alex! Hello Hello ..."))
}

