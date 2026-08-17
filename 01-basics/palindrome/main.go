package main

import (
	"fmt"
	"strings"
)

func palindrome(kata string) bool {
	runes := []rune(kata)
	var katabaru []rune

	for i := len(runes) - 1; i >= 0; i-- {
		huruf := runes[i]
		katabaru = append(katabaru, huruf)
	}

	return kata == string(katabaru)
}

func main() {
	kata := strings.ToLower("kasur rusak")
	fmt.Println(palindrome(kata))
}
