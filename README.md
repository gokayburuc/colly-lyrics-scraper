# Lyricstranslate Scraper using colly 

# Table of contents
- [Lyricstranslate Scraper using colly](#lyricstranslate-scraper-using-colly)
  - [Badges](#badges)
  - [Project Goals](#project-goals)
  - [Usage 🛂](#usage-🛂)
  - [Pre-Requirements 🔌](#pre-requirements-🔌)
    - [Colly](#colly)

## Badges 
![License](https://badgen.net/github/license/gokayburuc/colly-lyrics-scraper)
![License](https://badgen.net/github/commits/gokayburuc/colly-lyrics-scraper)

## Project Goals 

- [ ] Writing a Crawler to check all the lyric URLs with given conditions 
- [x] Parsing a Lyrics using GO 
- [x] Writing lyric into a .txt file 



## Usage 🛂


- Git clone this repo 
```bash 
git clone https://github.com/gokayburuc/colly-lyrics-scraper.git
```
- Chage directory into your folder 

```bash 
cd colly-lyrics-scraper/
```
- Build the project 

```bash 
go build .
```

- Run the project 
```bash 
./lyrics-scraper
```
- Enter Lyricstranslate.com URL and filename 
```bash 
New URL: 
# paste your lyricstranslate.com URL 
Filename: 
## enter proper filename 
```

> **DISCLAIMER:**: This scraper only downloads lyrics of a single URL for now. The bulk download option will also be added to the projects, but if not required, please prefer to read the text from the website.



## Pre-Requirements 🔌

### Colly 

> Colly is a Golang framework for building web scrapers. 

- [colly - Official Website](https://go-colly.org/) 
- [Gocolly - Github  ](https://github.com/gocolly/colly)


Installation : 
```bash
go get -u github.com/gocolly/colly/...
```


