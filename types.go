package main

import (
	"github.com/pimmytrousers/pastescraper/parse"
	log "github.com/sirupsen/logrus"
	"github.com/williballenthin/govt"
)

type Scraper struct {
	scrapingUrl    string
	pastesPerQuery int
	parser         *parse.Parser
	outputDir      string
	rawUrl         string
	debug          bool
	maxQueue       int
	logger         *log.Logger
	seenKeys       *keyQueue
	vtClient       *govt.Client
}

type Elastic struct {
	HTTPS    bool   `yaml:"https"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Index    string `yaml:"index"`
}

type config struct {
	OutputDir    string   `yaml:"outputdir"`
	VTKey        string   `yaml:"vtkey"`
	Debug        bool     `yaml:"debug"`
	MaxQueueSize int      `yaml:"maxqueuesize"`
	Elastic      Elastic  `yaml:"elastic"`
	Parsers      []string `yaml:"parsers"`
}

type PasteMetadata struct {
	ScrapeURL string `json:"scrape_url"`
	FullURL   string `json:"full_url"`
	Date      string `json:"date"`
	Key       string `json:"key"`
	Size      string `json:"size"`
	Expire    string `json:"expire"`
	Title     string `json:"title"`
	Syntax    string `json:"syntax"`
	User      string `json:"user"`
}
