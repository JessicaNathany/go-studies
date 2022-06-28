package main

import "fmt"

func Speak(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		fmt.Printf("%s: %s (interation %d)\n", person, text, i+1)
	}
}

func main() {
	Speak("Mariah", "Why do you don´t speak with me?", 3)
	Speak("Jhon", "I can only speak after you", 1)

	go Speak("Mariah", "hey...", 500)
	go Speak("Jhon", "Hi...", 500)
	fmt.Println("End!")
}
