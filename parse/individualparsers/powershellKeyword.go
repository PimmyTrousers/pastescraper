package individualparsers

import (
	"strings"
)

type PowershellKeyword struct{}

func (b PowershellKeyword) Match(content []byte) (bool, error) {
	// powershell contained within paste
	if strings.Contains(strings.ToLower(string(content)), "powershell") {
		return true, nil
	}

	return false, nil
}

func (b PowershellKeyword) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
