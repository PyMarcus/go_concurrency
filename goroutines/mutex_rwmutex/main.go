package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	global int = 0
	mutex  sync.Mutex
	lock   sync.RWMutex
	wait   sync.WaitGroup
)

func main() {
	for i := 0; i < 1000; i++ {
		go increasing()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Global value: ", global)

	readAndWrite()
}

func increasing() {
	mutex.Lock()
	defer mutex.Unlock()
	global++
}

func read() {
	fmt.Println("reading a large file...")
	time.Sleep(2 * time.Second)
	wait.Done()
}

func write() {
	fmt.Println("writen some lines in a file...")
	time.Sleep(1 * time.Second)
	wait.Done()
}

func readAndWrite() {
	lock.RLock()

	wait.Add(3)
	go read()
	go read()
	go write()
	lock.RUnlock()
	wait.Wait()
	fmt.Println("Finished!")
	time.Sleep(4 * time.Second)
}
