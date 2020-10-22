package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const sanskrit = "Sanskrit"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const sanskritHelloPrefix = "Namaste, "

// Hello will use the parameters to welcome the user
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case sanskrit:
		prefix = sanskritHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Name", "Language"))
}
