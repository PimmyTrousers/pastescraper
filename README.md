# pastescraper

A utility to scrape pastebin's incoming feed for known malware techniques.

You must have your ip whitelisted with pastebin for this to work.
[pastebin scraping docs](https://pastebin.com/doc_scraping_api)

## Usage

```go
// NOTE: This code does not check errors

// Create a parser object
parser, err := parse.New(strings.Split("base64MZHeader,base64ELFHeader", ","))

// Create a scraper
scraper, err := New(20, outputDir, parser)
err = scraper.start(time.Second * 10)
```
