package main

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/pimmytrousers/pastescraper/parse"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	SCRAPINGURL = "https://scrape.pastebin.com/api_scraping.php?limit=%d"
	RAWURL = "https://scrape.pastebin.com/api_scrape_item.php?i=%s"
)

type Scraper struct {
	scrapingUrl string
	pastesPerQuery int
	parser *parse.Parser
	outputDir string
	rawUrl string
	maxQueue int
	logger *log.Logger
	//TODO: use this to make sure you dont process keys you've already seen
	seenKeys  [100]string
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

type PasteStream []PasteMetadata

func (s *Scraper) GetRawPaste(key string) ([]byte, error) {
		client := http.Client{
			Timeout:       time.Second * 3,
		}

		rawUrlWithKey := fmt.Sprintf(s.rawUrl, key)
		resp, err := client.Get(rawUrlWithKey)
		if err != nil {
			return nil ,err
		}

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if len(buf) == 0 {
			return nil, errors.New("paste contents has length 0")
		}

		return buf, nil
}

func unmarshalPasteStream(data []byte) (PasteStream, error) {
	var r PasteStream
	err := json.Unmarshal(data, &r)
	return r, err
}

func (s *Scraper) getStreamChannel() (<-chan PasteMetadata, error) {
	out := make(chan PasteMetadata, s.maxQueue)
	client := http.Client{
		Timeout:       time.Second * 3,
	}

	resp, err := client.Get(fmt.Sprintf(s.scrapingUrl, s.pastesPerQuery))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	stream, err := unmarshalPasteStream(buf)

	//TODO: is this necessary?
	go func() {
		for _, pasteMetaData := range stream {
			out <- pasteMetaData
		}
		close(out)
	}()

	s.logger.Infof("added %d pastes to the queue", s.pastesPerQuery)
	return out, nil
}



func (s *Scraper) start(waitDuration time.Duration) error {
	for {
		stream, err := s.getStreamChannel()
		if err != nil {
			return err
		}

		//TODO: go routine this so the matching is done concurrently
		for pasteMetaData := range stream {
			pasteKey := pasteMetaData.Key
			pasteContent, err := s.GetRawPaste(pasteKey)
			if err != nil {
				s.logger.Warningf("failed to get paste contents for %s: %s", pasteKey, err)
				continue
			}

			matchedSig, err := s.parser.Match(pasteContent)
			if err != nil {
				s.logger.Warning("failed to get match content")
				continue
			}

			if matchedSig != "" {
				s.logger.WithFields(log.Fields{
					"signature match": matchedSig,
					"author": pasteMetaData.User,
					"size": pasteMetaData.Size,
					"title": pasteMetaData.Title,
					"full-url": pasteMetaData.FullURL,
					"key": pasteKey,
				}).Info("matched a paste")

				err = s.writePaste(matchedSig, pasteKey, pasteContent)
				if err != nil {
					s.logger.Warningf("failed to write paste %s: %s", pasteKey, err)
					continue
				}
			}
		}

		time.Sleep(waitDuration)
	}
}

func (s *Scraper) writePaste(key string, pasteKey string, content []byte) error {
	if _, err := os.Stat(s.outputDir); os.IsNotExist(err) {
		err = os.Mkdir(s.outputDir, 0755)
		if err != nil {
			return err
		}
	}

	parseSpecificPath := s.outputDir + "/" + key

	if _, err := os.Stat(parseSpecificPath); os.IsNotExist(err) {
		err = os.Mkdir(parseSpecificPath, 0755)
		if err != nil {
			return err
		}
	}

	err := ioutil.WriteFile(parseSpecificPath + "/" + pasteKey, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func New(maxQueueSize int, outputDir string, parser *parse.Parser) (*Scraper, error) {
	s := &Scraper{}

	s.outputDir = outputDir
	s.parser = parser
	s.maxQueue = maxQueueSize
	s.scrapingUrl = SCRAPINGURL
	s.rawUrl = RAWURL
	s.pastesPerQuery = 100
	s.logger = log.New()

	return s, nil
}
