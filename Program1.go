package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)
const saveDir = "./xkcd-comics"

func main() {
	//collector
	c := colly.NewCollector(
		colly.AllowedDomains("xkcd.com") // domain access
		colly.Async(true) //asyn for better performance
	)

	//creating directory for downloads, if one doesnt exist
	if error := os.MkdirAll(saveDir, os.ModePerm); != nil {
		log.Fatalf("Failed to create save directory: %v", error) //exits if creation fails
	}

	//what to do when image is found
	c.OnHTML("div#comic img", func(e *colly.HTMLElement){
		imgURL := e.Attr("scr") // gets scr attribute of img tag
		if !strings.HasPrefix(imgURL, "https:"){
			imgUrl = "https:" + imgURL //makes sure url is absolute
		}
		//download and save comic
		saveComic(imgURL)
	})

	

}




/*


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
