package individualparsers

import (
	"bytes"
)

type PowershellKeyword struct {}

func (b PowershellKeyword) Match(content []byte) (bool, error) {
	// powershell contained within paste
	if bytes.Contains(content, []byte("powershell")) {
		return true, nil
	}

	return false, nil
}

func (b PowershellKeyword) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
