package main

import (
	"fmt"

	//"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
	"github.com/cdipaolo/goml/text"
	//linear "github.com/cdipaolo/goml/perceptron"
)

func main() {
	// Define the text samples and corresponding labels
	samples := []string{
		"tell me a joke",
		"do not tell me a joke",
	}
	labels := []float64{1, 0}

	// Create the TF-IDF transformer
	tfidf := text.NewTFIDF(samples)

	// Transform the samples into TF-IDF vectors
	var trainData [][]float64
	for _, sample := range samples {
		vector, err := tfidf.Transform(sample)
		if err != nil {
			fmt.Println("Error transforming sample:", err)
			return
		}
		trainData = append(trainData, vector)
	}

	// Define and train the perceptron model
	model := linear.NewPerceptron(trainData, labels, nil)
	if err := model.Learn(); err != nil {
		fmt.Println("Error training model:", err)
		return
	}

	// Test the model with the given sentences
	for _, sample := range samples {
		vector, _ := tfidf.Transform(sample)
		prediction, _ := model.Predict(vector)
		fmt.Println("Predicted label for", sample, "is", prediction)
	}
}

