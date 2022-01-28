package main

import (
	"fmt"
)

//Languages
const (
	english = "english"
	spanish = "spanish"
	french  = "french"
)

//Greeting
const (
	englishGreeting = "Hello"
	spanishGreeting = "Hola"
	frenchGreeting  = "Bonjour"
)

func Hello(name string, language string) string {
	//handle empty string
	if name == "" {
		name = "World"
	}

	return greetingByLanguage(language) + ", " + name + "!"
}

func greetingByLanguage(language string) (greeting string) {
	switch language {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
