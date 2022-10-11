package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly/v2"
)

type Page struct {
	Heading    string
	HeroImage  string
	Paragraphs []string
}

func main() {
	// example usage: curl -s 'http://127.0.0.1:7171/?url=url'
	addr := ":7171"
	http.HandleFunc("/", scrapperHandler)
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func scrapAndDo(url string) *Page {
	page := new(Page)
	page.Paragraphs = make([]string, 0)

	c := colly.NewCollector()

	// Find header
	c.OnHTML(".entry-title", func(e *colly.HTMLElement) {
		page.Heading = e.Text
	})

	// get static hero image
	c.OnHTML(".separator", func(e *colly.HTMLElement) {
		page.HeroImage = e.ChildAttr("img", "src")
	})

	c.OnHTML(".post-body.entry-content span", func(e *colly.HTMLElement) {
		if len(e.Text) > 0 {
			page.Paragraphs = append(page.Paragraphs, e.Text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	fmt.Println("CRAWLING DONE:", page)

	return page
}
