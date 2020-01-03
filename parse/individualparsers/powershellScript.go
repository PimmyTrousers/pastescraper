package individualparsers

import (
	"bytes"
)

type PowershellScript struct {}

func (b PowershellScript) Match(content []byte) (bool, error) {
	// powershell invocation
	if bytes.Contains(content, []byte("Invoke-")) {
		return true, nil
	}

	return false, nil
}

func (b PowershellScript) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
