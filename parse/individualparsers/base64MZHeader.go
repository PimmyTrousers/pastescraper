package individualparsers

type Base64MZHeader struct {}

func (b Base64MZHeader) Match(content []byte) (bool, error) {
	// PE header base64 encoded
	if string(content[:4]) == "TVpQ" {
		return true, nil
	}

	return false, nil
}
