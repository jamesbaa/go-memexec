//go:build !linux
// +build !linux

package memexec

import (
	"fmt"
	"os"
	"runtime"
)

func open(b []byte, prefix string) (*os.File, error) {
	pattern := fmt.Sprintf("%s-", prefix)
	if runtime.GOOS == "windows" {
		pattern = fmt.Sprintf("%s-*.exe", prefix)
	}
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = clean(f)
		}
	}()
	if err = os.Chmod(f.Name(), 0o500); err != nil {
		return nil, err
	}
	if _, err = f.Write(b); err != nil {
		return nil, err
	}
	if err = f.Close(); err != nil {
		return nil, err
	}
	return f, nil
}

func clean(f *os.File) error {
	return os.Remove(f.Name())
}
