package individualparsers

import (
	"strings"
)

type PowershellKeyword struct{}

func (b PowershellKeyword) Match(content []byte) (bool, error) {
	// powershell contained within paste
	lowerContent := strings.ToLower(string(content))
	normalContent := strings.Replace(lowerContent, "^", "", -1)
	if strings.Contains(normalContent, "powershell") {
		return true, nil
	}

	return false, nil
}

func (b PowershellKeyword) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
