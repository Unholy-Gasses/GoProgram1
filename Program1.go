package main 

import (
	"github.com/gocolly/colly"
)

type comic struct {
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector() //creates a new instance of a Collector
	c.OnHTML("div=id"comic"", func(h *colly.HTMLElement)) //Any time CSS Selector ("") finds what its looking for, run func(), (h *colly.HTMLElement)
}