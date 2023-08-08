package main

import (
	"fmt"
    "os"
    "strings"

	"github.com/cdipaolo/sentiment"
)

func main() {

	// Define the text you want to analyze
	//text := "Golang is a fantastic programming language!"

    if len(os.Args) < 2 {
        fmt.Println("Please provide a text string to analyze.")
        return
    }

    //text := os.Args[1]
    // Get all the arguments starting from index 1 and concatenate them into a single string
    text := strings.Join(os.Args[1:], " ")

	// Initialize a new sentiment analysis model
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}


	// Get the analysis
	analysis := model.SentimentAnalysis(text, sentiment.English)

	// Print the result
	fmt.Println("Scores:", analysis.Score)
	//fmt.Println("Sentiment:", analysis.Desc)
}

