package main

import (
	"fmt"
	"time"
)

// calculate cousin numbers

func IsCousinNumber(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func Cousins(number int, channel chan int) {
	start := 2
	for i := 0; i < number; i++ {
		for cousin := start; ; cousin++ {
			channel <- cousin
			start = cousin + 1
			time.Sleep(time.Millisecond * 100)
			break
		}
	}
	close(channel)
}

func main() {
	c := make(chan int, 30)
	go Cousins(cap(c), c)
	for cousin := range c {
		fmt.Printf("%d ", cousin)
	}

	fmt.Println("End")
}
