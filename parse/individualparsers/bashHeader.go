package individualparsers

import (
	"strings"
)

type BashHeader struct{}

func (b BashHeader) Match(content []byte) (bool, error) {
	// bash header
	if len(content) < 11 {
		return false, nil
	}

	if strings.ToLower(string(content[:11])) == "#!/bin/bash" {
		return true, nil
	}

	return false, nil
}

func (b BashHeader) PostAction(content []byte) (string, error) {
	return "", nil
}

func (b BashHeader) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
