package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func (c *config) getConf(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, c)
}
