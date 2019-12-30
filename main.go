package main

import (
	"flag"
	"github.com/pimmytrousers/pastescraper/parse"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var (
	outputDir string
	flagParsers string
)

func init() {
	flag.StringVar(&outputDir, "outputdir", "./pastes", "output directory for pastes")
	flag.StringVar(&flagParsers, "parsers", "", "comma seperated list of parse to be used to saved pastes")
}

func main() {
	flag.Parse()
	if flagParsers == "" {
		flag.PrintDefaults()
	}

	parser, err := parse.New(strings.Split(flagParsers, ","))
	if err != nil {
		log.Fatalf("failed to initialize parser: %s", err)
	}

	scraper, err := New(20, outputDir, parser)
	if err != nil {
		log.Fatalf("failed to initialize scraper: %s", err)
	}

	err = scraper.start(time.Minute)
	if err != nil {
		log.Fatalf("failed to scrape: %s", err)
	}
}
