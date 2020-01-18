package individualparsers

import "encoding/base64"

type Base64MZHeader struct {}

func (b Base64MZHeader) Match(content []byte) (bool, error) {
	// PE header base64 encoded
	if len(content) < 4 {
		return false, nil
	}

	headerContents := string(content[:4])
	if headerContents == "TVpQ" || headerContents == "TVoA" || headerContents == "TVpB" || headerContents == "TVqA" || headerContents == "TVqQ" || headerContents == "TVro" {
		return true, nil
	}

	return false, nil
}

func (b Base64MZHeader) Normalize(content []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(content))
}
