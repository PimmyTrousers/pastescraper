package individualparsers

import (
	"bytes"
	"encoding/base64"
)

type Base64HighEntropy struct {}

func (b Base64HighEntropy) Match(content []byte) (bool, error) {
	// PE header base64 encoded
	base64DecodedContent, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		return false, err
	}

	entropy, err := entropy(bytes.NewReader(base64DecodedContent))
	if err != nil {
		return false, nil
	}

	if entropy > 7 {
		return true, nil
	}



	return false, nil
}

func (b Base64HighEntropy) Normalize(content []byte) ([]byte, error) {
	return content, nil
}

