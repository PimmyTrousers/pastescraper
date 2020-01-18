package individualparsers

import "encoding/base64"

type ReverseBase64ELFHeader struct {}

func (b ReverseBase64ELFHeader) Match(content []byte) (bool, error) {
	// ELF header base64 encoded
	reversedContent := reverse(string(content))
	if reversedContent[:5] == "f0VMR" {
		return true, nil
	}

	return false, nil
}

func (b ReverseBase64ELFHeader) Normalize(content []byte) ([]byte, error) {
	reversedContent := reverse(string(content))
	return base64.StdEncoding.DecodeString(string(reversedContent))
}
