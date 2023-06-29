package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	ninjas := []string{"NinjaA", "NinjaB"}
	smoke := make(chan bool)

	for _, ninja := range ninjas {
		go attack(ninja, smoke)
	}

	if <-smoke {
		fmt.Println("OK")
	}

}

func attack(ninja string, smoke chan bool) {
	fmt.Println("Throwing ninja start at ", ninja)
	time.Sleep(2 * time.Second)
	smoke <- false
}
