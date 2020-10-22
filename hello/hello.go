package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello will use the parameters to welcome the user
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Name"))
}
