package main

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func getConf(path string) (*config, error) {
	c := &config{}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func in(l []string, str string) bool {
	for _, s := range l {
		if s == str {
			return true
		}
	}

	return false
}

func maxCount(words map[string]int) string {
	max := 0
	topWord := ""
	for word, count := range words {
		if count > max {
			max = count
			topWord = word
		}
	}

	return topWord
}

func wordCountWithBlacklist(wordList []string, blacklist []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range wordList {
		word = strings.ToLower(word)
		if in(blacklist, word) {
			continue
		}
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}
