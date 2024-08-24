package main

import "fmt"

// const spanish = "Spanish"
// const french = "French"
// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "

// func Hello(name string, language string, greeting string) string {
// 	if name == "" {
// 		name = "World"
// 	}

// 	// if language == spanish {
// 	// 	return spanishHelloPrefix + name
// 	// }
// 	// if language == french {
// 	// 	return frenchHelloPrefix + name
// 	// }

// 	// return englishHelloPrefix + name

// 	// when we have lots of 'ifs' we should use a switch instead of the above
// 	prefix := englishHelloPrefix

// 	switch language {
// 	case "French":
// 		prefix = frenchHelloPrefix
// 	case "Spanish":
// 		prefix = spanishHelloPrefix
// 	}

// 	if greeting != "" {
// 		return greeting+", " + name
// 	}
// 	return prefix + name
// }

// ONE LAST REFACTOR WITH NOTES ABOUT WHAT WAS LEARNED HERE:
// Here we group our constants in a block instead of declaring them each on their own line. 
const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// we started this func with a lowercase since we don't want to export the inner workings of our
// algorithm to the world. Considered a 'private function' in Go
func greetingPrefix(language string) (prefix string) { // here we have a named return which creates
	// the var called prefix in the func. Its assigned to its zero default which is "" and we can then
	// just return instead of 'return prefix'. It will also display in Go Doc
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default: // our default case is none of the above work
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
