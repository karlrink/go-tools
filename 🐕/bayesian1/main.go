package main

import (
	"fmt"
	"strings"

	"github.com/jbrukh/bayesian"
)

const (
	JokeClass     bayesian.Class = "Tell me a joke"
	NoJokeClass   bayesian.Class = "Do not tell me a joke"
)

func main() {
	// Create a new classifier with the defined classes
	classifier := bayesian.NewClassifier(JokeClass, NoJokeClass)

	// Learn the two patterns
	classifier.Learn(strings.Fields("tell me a joke"), JokeClass)
	classifier.Learn(strings.Fields("do not tell me a joke"), NoJokeClass)

	// Convert the test sentences into a format the classifier can understand
	testSentence := "do not tell me a joke"
	words := strings.Fields(testSentence)

	// Perform classification
	scores, likely, _ := classifier.LogScores(words)

	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Most likely class for '%s': %s\n", testSentence, likely)
}

