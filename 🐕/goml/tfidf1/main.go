package main

import (
	"fmt"
	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
)

func main() {
	// Define the text samples and corresponding labels
	samples := []string{
		"tell me a joke",
		"do not tell me a joke",
	}
	labels := []float64{1, 2}

	// Initialize the vectorizer to convert text into TF-IDF vectors
	vectorizer := text.NewCountVectorizer()
	tfidfTransformer := text.NewTfidfTransformer()

	// Fit the vectorizer to the samples and transform them
	trainData, err := vectorizer.FitTransform(samples)
	if err != nil {
		fmt.Println("Error fitting/transforming samples:", err)
		return
	}

	// Transform the count vectors into TF-IDF vectors
	trainData, err = tfidfTransformer.FitTransform(trainData)
	if err != nil {
		fmt.Println("Error fitting/transforming TF-IDF:", err)
		return
	}

	// Initialize a Naive Bayes classifier to learn the difference between the sentences
	model := text.NewMultinomialNB(base.BatchGA, 0.00001, 0, trainData, labels)

	// Train the classifier
	if err = model.Learn(); err != nil {
		fmt.Println("Error training model:", err)
		return
	}

	// Test the model with an example sentence
	testSample := "tell me a joke"
	testData, _ := vectorizer.Transform([]string{testSample})
	testData, _ = tfidfTransformer.Transform(testData)

	prediction, _ := model.Predict(testData[0])
	fmt.Println("Predicted label for", testSample, "is", prediction) // Should output 1

	testSample = "do not tell me a joke"
	testData, _ = vectorizer.Transform([]string{testSample})
	testData, _ = tfidfTransformer.Transform(testData)

	prediction, _ = model.Predict(testData[0])
	fmt.Println("Predicted label for", testSample, "is", prediction) // Should output 2
}


