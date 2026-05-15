package main

import (
	"time"
)

func worker(workerId int, msg chan string) {
	for res := range msg {
		println("Worker", workerId, "got", res)
		time.Sleep(time.Second)
	}
}


func main() {

	msg := make(chan string, 1)

	msg <- "buffered"
	go worker(1, msg)

	msg <- "channel"
	go worker(1, msg)

	time.Sleep(time.Second * 2)
}
