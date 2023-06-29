package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("%v\n", time.Since(start))
	}()

	ninjas := []string{"ninjaA", "ninjaB", "ninjaC", "ninjaD"}

	for _, ninja := range ninjas {
		go attack(ninja)
	}

	time.Sleep(2 * time.Second) // a basic and poor form to synchronize the goroutines
}

func attack(ninjaName string) {
	fmt.Printf("Throwing ninja start at %s\n", ninjaName)
	time.Sleep(1 * time.Second)
}
