package individualparsers

type BashBang struct {}

func (b BashBang) Match(content []byte) (bool, error) {
	// bash header
	if string(content[:11]) == "#!/bin/bash" {
		return true, nil
	}

	return false, nil
}
