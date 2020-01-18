package individualparsers

import (
	"regexp"
)

type LargeHexBlob struct{}

func (b LargeHexBlob) Match(content []byte) (bool, error) {
	// large hex blobs 2A564D5A.....
	var re = regexp.MustCompile(`(?m)[2-9A-F]{200,}`)

	output := re.Find(content)
	if output != nil {
		return true, nil
	}
	return false, nil
}

func (b LargeHexBlob) Normalize(content []byte) ([]byte, error) {
	return content, nil
}
