package main

import "fmt"

func main() {
	channel := make(chan int)
	array := [5]int{1, 2, 3, 4, 5}
	go corroutine(array, channel)

	for i := 0; i < len(array); i++ {
		fmt.Println("response: ", <-channel)
	}

	go closeCorroutine(array, channel)

	for num := range channel {
		fmt.Println(num)
	}
}

// returns item to item
func corroutine(array [5]int, channel chan int) {
	for _, a := range array {
		channel <- a
	}
}

// close function allows says to program that it done!
func closeCorroutine(array [5]int, channel chan int) {
	for _, a := range array {
		channel <- a
	}

	close(channel)
}
