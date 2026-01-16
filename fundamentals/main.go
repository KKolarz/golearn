package main

import "fmt"

const englishHelloPrefix = "Hello, "
const helloSuffix = "!"

func Hello(name string) string {
	if name == "" {
		name = "stranger"
	}

	return englishHelloPrefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello("matt"))
}
