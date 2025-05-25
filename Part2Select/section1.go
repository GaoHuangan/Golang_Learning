package Part2Select

import (
	"time"
)

func SendPing(ch chan string) {
	time.Sleep(500 * time.Millisecond)
	ch <- "ping"
}

func SendPong(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "pong"
}