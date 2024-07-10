package scraper

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func SongText(url string) {

	c := colly.NewCollector(
		colly.AllowedDomains("lyricstranslate.com"),
		colly.CacheDir("lyrics_cache"),
	)

	// html üzerindeki veri
	// WARN: div id="song-body"
	c.OnHTML("div#song-body", func(h *colly.HTMLElement) {
		lyrics := h.Text // INFO: Text ifadesi () almaz
		fmt.Println(lyrics)

		filename := FilenameGenerator()

		WriteIntoFile(lyrics, filename)
	})

	// sistem request talebi
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Now:", r.URL)
	})

	c.Visit(url)

}

func FilenameGenerator() string {
	// INFO :filename
	var filename string
	fmt.Println("Filename:")
	fmt.Scan(&filename)
	return filename

}

func ErrCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}

func WriteIntoFile(data string, filename string) {

	fname := fmt.Sprintf("%s.txt", filename)
	f, err := os.Create(fname)

	defer f.Close()
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ErrCheck(err)

	x, err2 := f.WriteString(data)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	ErrCheck(err2)
	fmt.Print(x)
}

func GetURL() string {
	var URL string
	fmt.Println("New URL:")
	fmt.Scan(&URL)
	return URL
}

func GetSong() {
	// var url string = "https://lyricstranslate.com/en/Halid-Beslic-Miljacka-lyrics.html"
	url := GetURL()
	SongText(url)
}
