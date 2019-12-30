package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type PowershellKeyword struct {}

func (b PowershellKeyword) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// powershell contained within paste
	if bytes.Contains(buf, []byte("powershell")) {
		return true, nil
	}

	return false, nil
}
