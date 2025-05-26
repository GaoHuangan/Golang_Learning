package main

import (
	"fmt"
	"time"
)

// create a channel
func woker(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "job is finished"
}

func sayHallo() {
	fmt.Println("Hallo from goroutine")
}

func sendMsg(ch chan string) {
	ch <- "golang is best language"
}

func sendNumber(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i //sendNumber
	}
	close(ch) // close channel
}
