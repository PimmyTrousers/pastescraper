package individualparsers

import (
	"strings"
)

type PowershellScript struct{}

func (b PowershellScript) Match(content []byte) (bool, error) {
	// powershell invocation
	if strings.Contains(strings.ToLower(string(content)), "invoke-") {
		return true, nil
	}

	return false, nil
}

func (b PowershellScript) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
