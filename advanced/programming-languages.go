package main

import (
	"fmt"
	"strings"
)

type Languages struct {
	CSharp     string
	Python     string
	Go         string
	JavaScript string
	Java       string
	MatLab     string
	R          string
	Closure    string
}

func (language Languages) ProgrammingLanguagesInformation() string {
	// put information about programming languages
}

func ChooseLanguage() {
	// print options for select languages...
}

func PrintLanguage() {
	// call ProgrammingLanguagesInformation() ...
}

func main() {
	fmt.Println("Please, enter your name.")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s!", name)
	name = strings.TrimSpace(name)
}

// to be continued
