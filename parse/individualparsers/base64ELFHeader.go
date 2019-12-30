package individualparsers

type Base64ELFHeader struct {}

func (b Base64ELFHeader) Match(content []byte) (bool, error) {
	// ELF header base64 encoded
	if string(content[:5]) == "f0VMR" {
		return true, nil
	}

	return false, nil
}
