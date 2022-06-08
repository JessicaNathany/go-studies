package main

import "fmt"

// This simple condition control structure using IF/ELSE
func PrintResultNote(note float64) {
	if note >= 7 {
		fmt.Println("Aproved with note", note)
	} else {
		fmt.Println("Reproved with note", note)
	}
}

// continue...
func PrintResultSary(note float64) {

	if note >= 7 {
		fmt.Println("Aproved with note", note)
	} else {
		fmt.Println("Reproved with note", note)
	}
}

func main() {
	PrintResultNote(7.3)
	PrintResultNote(6.5)

}
