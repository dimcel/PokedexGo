package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	//var str []string

	str := strings.Fields(text)

	for i := range str {
		str[i] = strings.ToLower(str[i])
	}

	return str
}
