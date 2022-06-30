package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // I send the variable channel the value 1 (write)
	<-ch    // receveid the channel data (read)

	ch <- 2
	fmt.Println(<-ch)
}
