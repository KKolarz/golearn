package main

const spanish = "esp"
const polish = "pol"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const polishHelloPrefix = "Dzien dobry, "

const helloSuffix = "!"

func Hello(name string, lang string) string {
	if name == "" {
		name = "stranger"
	}

	prefix := greetingPrefix(lang)

	return prefix + name + helloSuffix
}

func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
}
