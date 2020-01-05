package individualparsers

import (
	"strings"
)

type PowershellFromBase64 struct {}

func (b PowershellFromBase64) Match(content []byte) (bool, error) {
	// powershell contained within paste
	if strings.Contains(strings.ToLower(string(content)), "frombase64string") {
		return true, nil
	}

	return false, nil
}

func (b PowershellFromBase64) Normalize(content []byte) ([]byte, error) {
	return content, nil
}

