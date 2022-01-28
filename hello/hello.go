package main

import "fmt"

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
	greeting := englishGreeting
	//handle language
	switch language {
	case english:
		greeting = englishGreeting
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	}
	return greeting + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
