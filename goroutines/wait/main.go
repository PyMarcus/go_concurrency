package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	words := []string{"Ninja1", "Ninja2"}
	rand.Seed(time.Now().UnixNano())

	c := make(chan string)
	go attack(words[rand.Intn(len(words))], c)

	select {
	case j := <-c:
		if j == "Ninja1" {
			fmt.Println("OK")
		} else {
			fmt.Println("Nop")
		}
	}

	signal := sync.WaitGroup{}
	signal.Add(2)
	// second
	go toWait("Hello", &signal)
	// first
	go toWait2("hi", &signal)
	signal.Wait()

	fmt.Println("All is finished!")
}

func attack(ninja string, c chan string) {
	fmt.Println(ninja)
	c <- ninja
}

func toWait(message string, signal *sync.WaitGroup) {
	fmt.Println("This is an operation ", message)
	signal.Done()
}

func toWait2(message string, signal *sync.WaitGroup) {
	fmt.Println("This is another operation ", message)
	signal.Done()
}
