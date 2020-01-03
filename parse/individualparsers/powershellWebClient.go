package individualparsers

import (
	"bytes"
)

type PowershellWebClient struct {}

func (b PowershellWebClient) Match(content []byte) (bool, error) {
	// powershell well client
	if bytes.Contains(content, []byte("Net.WebClient")) {
		return true, nil
	}

	return false, nil
}

func (b PowershellWebClient) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
