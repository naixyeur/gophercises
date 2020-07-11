package main

import "fmt"

func main() {
	var input string
	fmt.Print("Input: ")
	fmt.Scanf("%s", &input)

	b := []byte(input)

	fmt.Println(b)
	fmt.Println(input)
	wordCount := 1

	for _, c := range input {
		if c >= 'A' && c <= 'Z' {
			wordCount++
		}
	}

	fmt.Printf("word count: %v \n", wordCount)

}
