package individualparsers

import (
	"bytes"
)

type PowershellEncodedContent struct {}

func (b PowershellEncodedContent) Match(content []byte) (bool, error) {
	// powershell.exe -nop -wind hidden -Exec Bypass -noni -enc
	if bytes.Contains(content, []byte("-nop")) && bytes.Contains(content, []byte("-Exec Bypass")) && bytes.Contains(content, []byte("-Exec Bypass")) && bytes.Contains(content, []byte("-enc")) {
		return true, nil
	}

	return false, nil
}
