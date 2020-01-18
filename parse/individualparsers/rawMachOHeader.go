package individualparsers

import "bytes"

type RawMachOHeader struct{}

func (b RawMachOHeader) Match(content []byte) (bool, error) {
	// Raw Mach O header
	if len(content) < 4 {
		return false, nil
	}

	if bytes.Equal(content[:4], []byte{0xfe, 0xed, 0xfa, 0xcf}) {
		return true, nil
	}

	return false, nil
}

func (b RawMachOHeader) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
