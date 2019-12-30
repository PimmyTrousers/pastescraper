package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type PowershellWebClient struct {}

func (b PowershellWebClient) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// powershell well client
	if bytes.Contains(buf, []byte("Net.WebClient")) {
		return true, nil
	}

	return false, nil
}
