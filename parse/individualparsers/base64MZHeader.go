package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type Base64MZHeader struct {}

func (b Base64MZHeader) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// PE header base64 encoded
	if string(buf[:4]) == "TVpQ" {
		return true, nil
	}

	return false, nil
}
