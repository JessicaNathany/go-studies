package main

import (
	"fmt"
	"time"
)

func Routine(channel chan int) {
	time.Sleep(time.Second)

	channel <- 1 // block operation
}

func main() {
	c := make(chan int) //channel no buffer

	go Routine(c)

	fmt.Println(<-c) // block operation
	fmt.Println("Was read")
	fmt.Println(<-c) //deadlock
	fmt.Println("End")
}
