package individualparsers

import (
	"bytes"
	"io/ioutil"
)

type PythonSyscall struct {}

func (b PythonSyscall) Match(content []byte) (bool, error) {
	buf, err := ioutil.ReadAll(bytes.NewBuffer(content))
	if err != nil {
		return false, err
	}

	// python syscall
	if bytes.Contains(buf, []byte("os.system")) {
		return true, nil
	}

	return false, nil
}
