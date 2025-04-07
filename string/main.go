package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const str = "面"

	bytes := []byte(str)
	runes := []rune(str)

	fmt.Println("Bayt uzunluğu:", len(str))
	fmt.Println("Rune uzunluğu:", utf8.RuneCountInString(str))

	for _, byt := range bytes {
		fmt.Print(byt)
	}
	fmt.Println()
	for _, run := range runes {
		fmt.Print(run)
	}
	fmt.Println()
	fmt.Printf(string(runes))
	fmt.Println()
	fmt.Printf(string(bytes))
}
