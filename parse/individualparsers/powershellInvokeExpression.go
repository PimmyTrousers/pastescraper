package individualparsers

import (
	"strings"
)

type PowershellInvokeExpression struct{}

func (b PowershellInvokeExpression) Match(content []byte) (bool, error) {
	// powershell invoke expression used to execute certain commands in PS
	lowerContent := strings.ToLower(string(content))
	normalContent := strings.Replace(lowerContent, "^", "", -1)
	if strings.Contains(normalContent, "invoke-expression") {
		return true, nil
	}

	return false, nil
}

func (b PowershellInvokeExpression) Normalize(content []byte) (int, []byte, error) {
	return KeyNonActionable, content, nil
}
