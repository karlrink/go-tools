package main

import (
	"fmt"
	"strings"

	"github.com/jbrukh/bayesian"
)

func main() {
	// Define the classes
	JokeClass := "1"
	NoJokeClass := "0"

	// Create a new classifier with the defined classes
	classifier := bayesian.NewClassifierTfIdf(JokeClass, NoJokeClass, nil)

	// Training data
	trainingData := []struct {
		label   string
		sentence string
	}{
		{JokeClass, "tell me a joke"},
		{JokeClass, "please tell a joke"},
		{NoJokeClass, "do not tell me a joke"},
		{NoJokeClass, "please don't tell a joke"},
	}

	// Learn the training patterns
	for _, data := range trainingData {
		classifier.Learn(strings.Fields(data.sentence), data.label)
	}

	// Test sentences
	testSentences := []string{
		"tell me a joke",
		"please don't tell a joke",
	}

	for _, testSentence := range testSentences {
		words := strings.Fields(testSentence)
		_, likely, _ := classifier.LogScores(words)
		fmt.Printf("Most likely class for '%s': %s\n", testSentence, likely)
	}
}

