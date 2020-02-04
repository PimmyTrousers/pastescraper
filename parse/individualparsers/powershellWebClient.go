package individualparsers

import (
	"strings"
)

type PowershellWebClient struct{}

func (b PowershellWebClient) Match(content []byte) (bool, error) {
	// powershell well client
	lowerContent := strings.ToLower(string(content))
	normalContent := strings.Replace(lowerContent, "^", "", -1)
	if strings.Contains(normalContent, "net.webclient") {
		return true, nil
	}

	return false, nil
}

func (b PowershellWebClient) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
