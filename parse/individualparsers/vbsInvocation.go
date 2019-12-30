package individualparsers

import (
	"bytes"
)

type VbsInvocation struct {}

func (b VbsInvocation) Match(content []byte) (bool, error) {
	// visual basic shell invocation
	if bytes.Contains(content, []byte("wscript.shell")) {
		return true, nil
	}

	return false, nil
}
