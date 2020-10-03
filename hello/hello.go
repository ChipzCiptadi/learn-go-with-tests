package main

import "fmt"

const englishHelloPrefix = "Hello"
const bahasaHelloPrefix = "Halo"
const batakHelloPrefix = "Horas"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Bahasa":
		prefix = bahasaHelloPrefix
	case "Batak":
		prefix = batakHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("World", ""))
}
