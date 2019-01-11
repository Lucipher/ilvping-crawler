package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Before making a request put the URL with
	// the key of "url" into the context of the request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	// After making a request get "url" from
	// the context of the request
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit("http://www.ilvping.com/view/Index/index/t%3D1%26ac%3D1.html")

}
