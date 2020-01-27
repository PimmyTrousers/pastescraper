package individualparsers

import (
	"encoding/base64"
)

type ReverseBase64ELFHeader struct{}

func (b ReverseBase64ELFHeader) Match(content []byte) (bool, error) {
	// ELF header base64 encoded
	reversedContent := reverse(string(content))
	if len(reversedContent) < 5 {
		return false, nil
	}

	if reversedContent[:5] == "f0VMR" {
		return true, nil
	}

	return false, nil
}

func (b ReverseBase64ELFHeader) Normalize(content []byte) (int, []byte, error) {
	reversedContent := reverse(string(content))
	content, err := base64.StdEncoding.DecodeString(reversedContent)
	if err != nil {
		return 0, nil, err
	}

	return KeyRawExecutable, content, err
}
