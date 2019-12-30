package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type BashBang struct {}

func (b BashBang) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// bash header
	if string(buf[:11]) == "#!/bin/bash" {
		return true, nil
	}

	return false, nil
}
