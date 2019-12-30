package individualparsers

import (
	"bytes"
)

type PythonSyscall struct {}

func (b PythonSyscall) Match(content []byte) (bool, error) {
	// python syscall
	if bytes.Contains(content, []byte("os.system")) {
		return true, nil
	}

	return false, nil
}
