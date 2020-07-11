package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	r := strings.NewReader("some io.Reader stream to be read \n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T \n", os.Stdout)

}
