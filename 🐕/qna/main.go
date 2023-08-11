package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// Define a callback to process each visited HTML element
	c.OnHTML("p", func(e *colly.HTMLElement) {
		// Use a regex to match questions (ending with a question mark)
		re := regexp.MustCompile(`[A-Z][^.!?]*\?`)
		questions := re.FindAllString(e.Text, -1)
		for _, question := range questions {
			fmt.Println("Found question:", question)
		}
	})

	// Visit the Wikipedia page
	err := c.Visit("https://en.wikipedia.org/wiki/Internet")
	if err != nil {
		log.Fatal(err)
	}
}

