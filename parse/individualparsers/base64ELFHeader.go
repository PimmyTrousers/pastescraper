package individualparsers

import (
	"encoding/base64"
)

type Base64ELFHeader struct{}

func (b Base64ELFHeader) Match(content []byte) (bool, error) {
	// ELF header base64 encoded
	if len(content) < 5 {
		return false, nil
	}

	if string(content[:5]) == "f0VMR" {
		return true, nil
	}

	return false, nil
}

func (b Base64ELFHeader) Normalize(content []byte) (int, []byte, error) {
	content, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		return 0, nil, err
	}

	return KeyRawExecutable, content, err
}
