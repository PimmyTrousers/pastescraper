package individualparsers

import (
	"strings"
)

type PowershellScript struct{}

func (b PowershellScript) Match(content []byte) (bool, error) {
	// powershell invocation
	lowerContent := strings.ToLower(string(content))
	normalContent := strings.Replace(lowerContent, "^", "", -1)
	if strings.Contains(normalContent, "invoke-") {
		return true, nil
	}

	return false, nil
}

func (b PowershellScript) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
