package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

var (
	n         = flag.Int("n", 5, "Maximum number of parallel requests")
	root      = flag.String("root", "https://en.wikipedia.org/wiki/Money_Heist", "started page")
	recursion = flag.Int("r", 15, "recursion depth")
	useragent = flag.String("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36", "User-Agent")
)

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(*recursion),
		colly.UserAgent(*useragent),
	)
	flag.Parse()

	URL := *root
	if URL == "" {
		return
	}

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: *n})
	links := make(map[string]int)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			links[link]++
		}
	})

	c.Visit(URL)

	for link, count := range links {
		fmt.Println(count, link)
	}

	log.Println("listening on", ":9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
