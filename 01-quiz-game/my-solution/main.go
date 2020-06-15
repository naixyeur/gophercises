package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

var problems []problem
var score int
var doneCh = make(chan struct{})

var startTime time.Time
var timeout time.Time

func init() {
	rand.Seed(time.Now().UnixNano())

}

func main() {

	// ==================================================
	// Read File
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		problems = append(problems, problem{record[0], record[1]})

	}

	// ==================================================

	scanner := bufio.NewScanner(os.Stdin)
	// var timeoutCh <-chan time.Time
	for {
		// clear.Clear()
		fmt.Println("press enter to start...")
		scanner.Scan()

		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
		startTime = time.Now()
		timeout = time.Now().Add(10 * time.Second)
		timeoutCh := time.After(10 * time.Second)
		// fmt.Printf("%T  %T \n", timeoutCh, to)

		go quiz()

		select {
		case <-doneCh:
			fmt.Println("\nfinished!")

		case <-timeoutCh:
			fmt.Println("\ntimes up!")

		}
		fmt.Printf("\nyour score: %v \n", score)
	}

}

func quiz() {
	for _, p := range problems {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("what %v, sir? ", p.question)
		scanner.Scan()
		input := scanner.Text()

		if input == p.answer {
			score++
			fmt.Printf("correct     \t")

		} else {
			fmt.Printf("wrong     \t")
		}
		fmt.Printf("time left: %v \n", timeout.Sub(time.Now()))

	}
	doneCh <- struct{}{}
}
