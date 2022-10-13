package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly/v2"
)

type PageMeta struct {
	Title       string
	Description string
}

type Page struct {
	Heading    string
	Meta       PageMeta
	HeroImage  string
	Paragraphs []string
}

func main() {
	// example usage: curl -s 'http://127.0.0.1:7171/?url=url'
	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}

	port = fmt.Sprintf(":%s", port)

	http.HandleFunc("/", scrapperHandler)
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func scrapAndDo(url string) *Page {
	page := new(Page)
	page.Paragraphs = make([]string, 0)

	c := colly.NewCollector()

	// Find header
	c.OnHTML(".entry-title", func(e *colly.HTMLElement) {
		page.Heading = e.Text
		page.Meta.Title = e.Text
	})

	// get static hero image
	c.OnHTML(".separator", func(e *colly.HTMLElement) {
		page.HeroImage = e.ChildAttr("img", "src")
	})

	c.OnHTML(".post-body.entry-content span", func(e *colly.HTMLElement) {
		if len(e.Text) > 0 {
			// add meta info
			page.Meta.Description = e.Text
			page.Paragraphs = append(page.Paragraphs, e.Text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	fmt.Println("generated pages: ", len(page.Paragraphs))

	return page
}
