package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please, enter your name.")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s!", name)
	name = strings.TrimSpace(name)

}

// to be continued
