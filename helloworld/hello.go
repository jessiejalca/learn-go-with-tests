package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	german  = "German"

	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
	germanHello  = "Hallo, "
)

func Hello(name string, language string) string {
	// Default to world if no name is provided
	if name == "" {
		name = "World"
	}
	// Greet in different languages
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	case german:
		prefix = germanHello
	default:
		prefix = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
