package individualparsers

import (
	"bytes"
)

type PowershellFromBase64 struct {}

func (b PowershellFromBase64) Match(content []byte) (bool, error) {
	// powershell contained within paste
	if bytes.Contains(content, []byte("FromBase64String")) {
		return true, nil
	}

	return false, nil
}

