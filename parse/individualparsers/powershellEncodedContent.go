package individualparsers

import (
	"strings"
)

type PowershellEncodedContent struct{}

func (b PowershellEncodedContent) Match(content []byte) (bool, error) {
	// powershell.exe -nop -wind hidden -Exec Bypass -noni -enc
	if strings.Contains(strings.ToLower(string(content)), "-nop") &&
		strings.Contains(strings.ToLower(string(content)), "-exec bypass") &&
		strings.Contains(strings.ToLower(string(content)), "-enc") {
		return true, nil
	}

	return false, nil
}

func (b PowershellEncodedContent) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
