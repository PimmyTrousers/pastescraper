package individualparsers

import (
	"strings"
)

type PowershellInvokeExpression struct{}

func (b PowershellInvokeExpression) Match(content []byte) (bool, error) {
	// powershell invoke expression used to execute certain commands in PS
	if strings.Contains(strings.ToLower(string(content)), "invoke-expression") {
		return true, nil
	}

	return false, nil
}

func (b PowershellInvokeExpression) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
