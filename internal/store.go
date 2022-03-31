package internal

import (
	"io"
	"os"
)

// Store stores v to writer and returns an error if failed.
func Store(writer io.Writer, v interface{}) error {
	entries, err := mapFromStruct(v)
	if err != nil {
		return err
	}

	target, err := encodeEntries(entries)
	if err != nil {
		return err
	}

	_, err = io.WriteString(writer, target)
	return err
}

// StoreFile stores v to target file and returns an error if failed.
func StoreFile(target string, v interface{}) error {
	file, err := os.OpenFile(target, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return Store(file, v)
}
