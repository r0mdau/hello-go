package hello

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(langage string) string {
	switch langage {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	}
	return englishHelloPrefix
}

// Hello returns a greeting
func Hello(name string, langage string) string {
	if len(name) == 0 {
		name = "World"
	}
	return greetingPrefix(langage) + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
