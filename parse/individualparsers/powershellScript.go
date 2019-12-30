package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type PowershellScript struct {}

func (b PowershellScript) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// powershell invocation
	if bytes.Contains(buf, []byte("Invoke-")) {
		return true, nil
	}

	return false, nil
}
