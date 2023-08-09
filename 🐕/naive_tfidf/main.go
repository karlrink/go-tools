package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	samples := []string{
		"tell me a joke",
		"do not tell me a joke",
	}

	// Compute the term frequencies (TF)
	var termFrequencies []map[string]float64
	for _, sample := range samples {
		tf := make(map[string]float64)
		words := strings.Fields(sample)
		for _, word := range words {
			tf[word]++
		}
		for word := range tf {
			tf[word] /= float64(len(words))
		}
		termFrequencies = append(termFrequencies, tf)
	}

	// Compute the inverse document frequencies (IDF)
	idf := make(map[string]float64)
	for _, tf := range termFrequencies {
		for word := range tf {
			idf[word]++
		}
	}
	for word := range idf {
		idf[word] = math.Log(float64(len(samples)) / idf[word])
	}

	// Compute the TF-IDF vectors
	var tfidfVectors [][]float64
	for _, tf := range termFrequencies {
		var vector []float64
		for word := range idf {
			vector = append(vector, tf[word]*idf[word])
		}
		tfidfVectors = append(tfidfVectors, vector)
	}

	fmt.Println("TF-IDF vectors:", tfidfVectors)
}

