package main

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, logs := range log {
		switch logs {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	logs := []rune(log)
	for i, val := range logs {
		if val == oldRune {
			logs[i] = newRune
		}
	}
	return string(logs)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit

}

func main() {
	fmt.Println(WithinLimit("hello❗", 6))    //6
	fmt.Println(WithinLimit("exercism❗", 9)) //9
	fmt.Println(WithinLimit("exercism🔍", 9)) //9
}
