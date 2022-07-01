package main

import (
	"fmt"
	"time"
)

func Talk(person string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speaking: %d", person, i)
		}
	}()
	return c
}

func join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := join(Talk("Jhon"), Talk("Mariah"))
	println(<-c, <-c)
	println(<-c, <-c)
	println(<-c, <-c)
}
