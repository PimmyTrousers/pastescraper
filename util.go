package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
