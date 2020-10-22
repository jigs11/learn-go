package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello will use the parameters to welcome the user
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}

	if language == french {
		return frenchHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Name", "Language"))
}
