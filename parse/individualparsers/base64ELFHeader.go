package individualparsers

import "encoding/base64"

type Base64ELFHeader struct {}

func (b Base64ELFHeader) Match(content []byte) (bool, error) {
	// ELF header base64 encoded
	if string(content[:5]) == "f0VMR" {
		return true, nil
	}

	return false, nil
}

func (b Base64ELFHeader) Normalize(content []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(content))
}
