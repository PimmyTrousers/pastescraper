package individualparsers

import "bytes"

type Raw64MZHeader struct {}

func (b Raw64MZHeader) Match(content []byte) (bool, error) {
	// Raw MZ header
	if bytes.Equal(content[:3], []byte{0x4d, 0x5a, 0x90}) {
		return true, nil
	}

	return false, nil
}

func (b Raw64MZHeader) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
