package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type VbsInvocation struct {}

func (b VbsInvocation) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// visual basic shell invocation
	if bytes.Contains(buf, []byte("wscript.shell")) {
		return true, nil
	}

	return false, nil
}
