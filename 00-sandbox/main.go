package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// 	break
	// }

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

}
