package individualparsers

import (
	"strings"
)

type PowershellFromBase64 struct{}

func (b PowershellFromBase64) Match(content []byte) (bool, error) {
	// powershell contained within paste
	lowerContent := strings.ToLower(string(content))
	normalContent := strings.Replace(lowerContent, "^", "", -1)
	if strings.Contains(normalContent, "frombase64string") {
		return true, nil
	}

	return false, nil
}

func (b PowershellFromBase64) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
