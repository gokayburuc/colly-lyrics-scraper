package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func Crawler() {

	//create a collector
	c := colly.NewCollector(
		colly.AllowedDomains("lyricstranslate.com"),
		colly.IgnoreRobotsTxt(),
		colly.CacheDir("./lyrics_cache"),
		colly.DisallowedDomains("google.com", "facebook.com", "twitter.com"),
		// colly.CacheDir("./lyricstranslate_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div", func(h *colly.HTMLElement) {
		class := h.Attr("class")
		if class == "stt" {
			link := h.ChildAttr("a", "href")
			text := h.ChildText("a")
			newURL := h.Request.AbsoluteURL(link)
			fmt.Println("Song:", newURL)
			fmt.Println("Text:", text)
			if len(text) > 0 {
				h.Request.Visit(link)
			}
		}
	})

	err := c.Visit("https://lyricstranslate.com/en/language/bosnian-lyrics")

	if err != nil {
		log.Fatal(err)
	}
}
