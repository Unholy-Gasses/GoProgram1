package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Corrected URL string
	scrapeUrl := "https://xkcd.com/"

	// Create a new Colly collector
	c := colly.NewCollector(
		colly.AllowedDomains("xkcd.com"), // Allowed domain
	)

	// Define what to do when a matching HTML element is found
	c.OnHTML("div#comic", func(h *colly.HTMLElement) {
		fmt.Printf("Comic Text:", h.Text)
	})
	//A function that runs everytime a request goes through
	c.OnRequest((func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("Visiting %s", r.URL))
	}))
	//Prints error if it occurs
	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})
	// Visit URL
	c.Visit(scrapeUrl)
}
