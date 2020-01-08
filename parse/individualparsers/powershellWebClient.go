package individualparsers

import (
	"strings"
)

type PowershellWebClient struct {}

func (b PowershellWebClient) Match(content []byte) (bool, error) {
	// powershell well client
	if strings.Contains(strings.ToLower(string(content)), "net.webclient") {
		return true, nil
	}

	return false, nil
}

func (b PowershellWebClient) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
