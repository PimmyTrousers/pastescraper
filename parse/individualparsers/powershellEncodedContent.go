package individualparsers

import (
	"strings"
)

type PowershellEncodedContent struct{}

func (b PowershellEncodedContent) Match(content []byte) (bool, error) {
	// powershell.exe -nop -wind hidden -Exec Bypass -noni -enc
	lowerContent := strings.ToLower(string(content))
	normalContent := strings.Replace(lowerContent, "^", "", -1)
	if strings.Contains(normalContent, "-nop") &&
		strings.Contains(normalContent, "-exec bypass") &&
		strings.Contains(normalContent, "-enc") {
		return true, nil
	}

	return false, nil
}

func (b PowershellEncodedContent) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
