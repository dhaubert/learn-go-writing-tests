package main

import "fmt"

const spanish = "es_ES"
const french = "fr_FR"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefixLanguage(language) + name
}

func greetingPrefixLanguage(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	// based on https://dev.to/quii/learn-go-by-writing-tests--m63
	fmt.Println(Hello("world", ""));
}



