package individualparsers

import "strings"

type BashHeader struct {}

func (b BashHeader) Match(content []byte) (bool, error) {
	// bash header
	if strings.ToLower(string(content[:11])) == "#!/bin/bash" {
		return true, nil
	}

	return false, nil
}

func (b BashHeader) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
