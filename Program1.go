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

/*package main // Defines the main package, the entry point for the Go program.

import (
	"fmt"       // Provides functions for formatted I/O, like printing to the console.
	"log"       // Provides logging functionality for error and info messages.
	"net/http"  // Allows making HTTP requests (used to download images).
	"os"        // Handles file and directory operations (e.g., creating directories).
	"path/filepath" // Handles file path manipulations in a cross-platform way.
	"strings"   // Provides string manipulation utilities.
	"github.com/gocolly/colly" // The Colly library for web scraping.
)

// Directory where the comics will be saved.
const saveDir = "./xkcd-comics"

func main() {
	// Create a new Colly collector to manage scraping tasks.
	c := colly.NewCollector(
		colly.AllowedDomains("xkcd.com"), // Restrict scraping to the XKCD domain.
		colly.Async(true),                // Enable asynchronous scraping for better performance.
	)

	// Create the directory for saving comics if it doesn't already exist.
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create save directory: %v", err) // Exit if directory creation fails.
	}

	// Define what to do when an image is found inside the `div#comic` element.
	c.OnHTML("div#comic img", func(e *colly.HTMLElement) {
		imgURL := e.Attr("src") // Get the 'src' attribute of the <img> tag.
		if !strings.HasPrefix(imgURL, "https:") {
			imgURL = "https:" + imgURL // Ensure the URL is absolute by adding "https:".
		}

		// Download and save the comic image.
		saveComic(imgURL)
	})

	// Define what to do when the "Previous" button is found.
	c.OnHTML("a[rel='prev']", func(e *colly.HTMLElement) {
		prevURL := e.Request.AbsoluteURL(e.Attr("href")) // Get the absolute URL of the "Previous" comic link.
		fmt.Println("Next comic URL to visit:", prevURL) // Log the URL of the next comic to scrape.
		e.Request.Visit(prevURL) // Visit the previous comic URL recursively.
	})

	// Define what to do if an error occurs during scraping.
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error: %v (status code: %d)", err, r.StatusCode) // Log the error and HTTP status code.
	})

	// Start scraping from the latest comic on the XKCD homepage.
	fmt.Println("Starting scrape from the latest comic...")
	c.Visit("https://xkcd.com") // Begin the scraping process.

	// Wait for all asynchronous tasks to finish before exiting.
	c.Wait()
}

// Downloads an image from the provided URL and saves it locally.
func saveComic(imgURL string) {
	// Extract the file name from the URL.
	segments := strings.Split(imgURL, "/") // Split the URL by "/" to isolate the file name.
	fileName := segments[len(segments)-1] // The last segment is the file name.
	filePath := filepath.Join(saveDir, fileName) // Combine the save directory with the file name.

	// Skip downloading if the file already exists.
	if _, err := os.Stat(filePath); err == nil { // Check if the file exists.
		fmt.Printf("Comic %s already exists, skipping...\n", fileName) // Log a message if it exists.
		return
	}

	// Download the image using an HTTP GET request.
	resp, err := http.Get(imgURL) // Fetch the image from the URL.
	if err != nil {
		log.Printf("Failed to download %s: %v", imgURL, err) // Log an error if the download fails.
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed after the function ends.

	// Create the file locally for saving the image.
	out, err := os.Create(filePath) // Open a file for writing at the specified path.
	if err != nil {
		log.Printf("Failed to create file %s: %v", filePath, err) // Log an error if file creation fails.
		return
	}
	defer out.Close() // Ensure the file is properly closed after writing.

	// Copy the image data from the response body to the file.
	_, err = io.Copy(out, resp.Body) // Save the image to the file.
	if err != nil {
		log.Printf("Failed to save %s: %v", filePath, err) // Log an error if saving fails.
		return
	}

	// Log success message once the comic is saved.
	fmt.Printf("Successfully saved comic: %s\n", filePath)
}
*/
