package individualparsers

import (
	"strings"
)

type VbsInvocation struct {}

func (b VbsInvocation) Match(content []byte) (bool, error) {
	// visual basic shell invocation
	if strings.Contains(strings.ToLower(string(content)), "wscript.shell") {
		return true, nil
	}

	return false, nil
}

func (b VbsInvocation) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
