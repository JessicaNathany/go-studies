package main

import (
	"fmt"
	"time"
)

func Routine(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
	fmt.Println("Executed")
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	ch := make(chan int, 3)
	go Routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
