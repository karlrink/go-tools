package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{
		"tell me a joke",
		"do not tell me a joke",
	}

	// Create a bag-of-words model
	bow := make(map[string]int)
	for _, sentence := range sentences {
		for _, word := range strings.Fields(sentence) {
			bow[word]++
		}
	}

	// Check the presence of "not" in the sentences
	for _, sentence := range sentences {
		_, containsNot := bow["not"]
		if containsNot {
			fmt.Println(sentence, "contains 'not'")
		} else {
			fmt.Println(sentence, "does not contain 'not'")
		}
	}
}

// bag-of-words approach (without TF-IDF) to classify
