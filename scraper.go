package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pimmytrousers/pastescraper/parse"
	log "github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	SCRAPINGURL = "https://scrape.pastebin.com/api_scraping.php?limit=%d"
	RAWURL = "https://scrape.pastebin.com/api_scrape_item.php?i=%s"
)



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

	if s.debug {
		s.logger.WithFields(log.Fields{
			"rawContentsURL":        rawUrlWithKey,
			"key":                   key,
		}).Debug("got raw contents of paste")
	}

	return buf, nil
}

func unmarshalPasteStream(data []byte) ([]PasteMetadata, error) {
	var r []PasteMetadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (s *Scraper) getStreamChannel() ([]PasteMetadata, error) {
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
	if len(stream) == 0 {
		return nil, errors.New("unable to acquire a paste stream - most likely due to unwhitelisted IP")
	}

	if s.debug {
		s.logger.WithFields(log.Fields{
			"pastesAdded":        len(stream),
		}).Debug("acquired pastes from pastebin API")
	}

	return stream, nil
}

func (s *Scraper) start(ctx context.Context, waitDuration time.Duration) error {
	for {

		var stream []PasteMetadata
		var err error
		for i := 0; i < 5; i++ {
			stream, err = s.getStreamChannel()
			if err != nil {
				s.logger.WithFields(log.Fields{
					"error": 		   err,
				}).Warning("unable to get paste stream, trying again")
			}

			if stream != nil {
				break
			}

			time.Sleep(time.Second * 3)
		}

		if stream != nil {
			s.logger.WithFields(log.Fields{
				"error": 		   err,
			}).Warning("unable to get paste stream")
		}

		for _, pasteMetaData := range stream {
			//start a goroutine to match each paste in the stream
			go func(metadata PasteMetadata) {
				pasteKey := metadata.Key
				if s.seenKeys.doesExist(pasteKey) {
					if s.debug {
						s.logger.WithFields(log.Fields{
							"full-url":        metadata.FullURL,
							"key":             pasteKey,
						}).Debug("already parsed paste")
					}
					return
				}

				pasteContent, err := s.GetRawPaste(pasteKey)
				if err != nil {
					s.logger.WithFields(log.Fields{
						"full-url":        metadata.FullURL,
						"key":             pasteKey,
						"error": 		   err,
					}).Warning("unable to get raw paste")
					return
				}

				matchedSig, normalizedContent, err := s.parser.MatchAndNormalize(pasteContent)
				if err != nil {
					s.logger.WithFields(log.Fields{
						"full-url":        metadata.FullURL,
						"key":             pasteKey,
						"error": 		   err,
					}).Warning("unable to match against parsers")
					return
				}

				s.seenKeys.add(pasteKey)

				size, err := strconv.Atoi(metadata.Size)
				if err != nil {
					s.logger.WithFields(log.Fields{
						"error": 		   err,
					}).Warning("invalid size")
					return
				}

				if matchedSig != "" {
					s.logger.WithFields(log.Fields{
						"signature match": matchedSig,
						"author":          metadata.User,
						"size":            size,
						"title":           metadata.Title,
						"full-url":        metadata.FullURL,
						"key":             pasteKey,
					}).Info("matched a paste")

					err = s.writePaste(matchedSig, pasteKey, normalizedContent)
					if err != nil {
						s.logger.WithFields(log.Fields{
							"full-url":        metadata.FullURL,
							"key":             pasteKey,
							"error": 		   err,
						}).Warning("unable to write paste")
						return
					}
				} else {
					if s.debug {
						s.logger.WithFields(log.Fields{
							"author":          metadata.User,
							"size":            size,
							"title":           metadata.Title,
							"full-url":        metadata.FullURL,
							"key":             pasteKey,
						}).Info("unable to match paste")
					}
				}
			}(pasteMetaData)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
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

	if s.debug {
		s.logger.WithFields(log.Fields{
			"pastekey":                   pasteKey,
			"paste location":             parseSpecificPath + "/" + pasteKey,
		}).Debug("wrote paste contents to disk")
	}

	err := ioutil.WriteFile(parseSpecificPath + "/" + pasteKey, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func New(c *config, parser *parse.Parser) (*Scraper, error) {
	s := &Scraper{}

	s.outputDir = c.OutputDir
	s.debug = c.Debug
	s.parser = parser
	s.seenKeys = newKeyQueue(c.MaxQueueSize * 10)
	s.maxQueue = c.MaxQueueSize
	s.scrapingUrl = SCRAPINGURL
	s.rawUrl = RAWURL
	s.pastesPerQuery = c.MaxQueueSize
	s.logger = log.New()

	if s.debug {
		s.logger.SetLevel(log.DebugLevel)
	}

	if c.Elastic.Host != "" {
		var proto string

		if c.Elastic.HTTPS {
			proto = "https://"
		} else {
			proto = "http://"
		}

		var err error
		var client *elastic.Client

		url := proto + c.Elastic.Host + ":" + strconv.Itoa(c.Elastic.Port)

		if c.Elastic.Password != "" && c.Elastic.Username != "" {
			client, err = elastic.NewSimpleClient(elastic.SetURL(url), elastic.SetBasicAuth(c.Elastic.Username, c.Elastic.Password))
		} else {
			client, err = elastic.NewSimpleClient(elastic.SetURL(url))
		}

		if err != nil {
			s.logger.Fatal(err)
		}


		hook, err := elogrus.NewAsyncElasticHook(client, c.Elastic.Host, log.DebugLevel, c.Elastic.Index)
		if err != nil {
			s.logger.Fatal(err)
		}

		s.logger.WithFields(log.Fields{
			"url":                   url,
		}).Debug("connected to ELK instance")
		s.logger.Hooks.Add(hook)
	}

	return s, nil
}