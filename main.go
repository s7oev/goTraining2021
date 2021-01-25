package main

import "fmt"

const (
	Bulgarian = "bulgarian"
	English = "english"
	Romanian = "romanian"
	German = "german"

	BulgarianGreeting = "здравей"
	EnglishGreeting = "Hello"
	RomanianGreeting = "Salut"
	GermanGreeting = "Hallo"
)

func main() {
	fmt.Println(Greeting("Waldemar", ""))
}

func Greeting(name, language string) string {
	var greeting = "Hello"

	if name == "" {
		name = "World"
	}

	switch language {
	case Bulgarian:
		greeting = BulgarianGreeting
	case Romanian:
		greeting = RomanianGreeting
	case German:
		greeting = RomanianGreeting
	default:
		greeting = EnglishGreeting
	}

	return fmt.Sprintf("%s %s", greeting, name)
}