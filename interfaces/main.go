package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetingEnglish(eb)
	printGreetingSpanish(sb)

}

func printGreetingEnglish(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
func printGreetingSpanish(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic specific to the english bot
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola !"
}
