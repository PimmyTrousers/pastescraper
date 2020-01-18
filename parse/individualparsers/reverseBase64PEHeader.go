package individualparsers

import "encoding/base64"

type ReverseBase64MZHeader struct{}

func (b ReverseBase64MZHeader) Match(content []byte) (bool, error) {
	// PE header base64 encoded
	normalizedContent := reverse(string(content))
	if len(normalizedContent) < 4 {
		return false, nil
	}

	headerContents := normalizedContent[:4]

	if headerContents == "TVpQ" || headerContents == "TVoA" || headerContents == "TVpB" || headerContents == "TVqA" || headerContents == "TVqQ" || headerContents == "TVro" {
		return true, nil
	}

	return false, nil
}

func (b ReverseBase64MZHeader) Normalize(content []byte) ([]byte, error) {
	normalizedContent := reverse(string(content))
	return base64.StdEncoding.DecodeString(normalizedContent)
}
