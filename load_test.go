// Author: fish
// Email: yezijie@bilibili.com
// Created at 2020/11/13 11:36

package props

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

// prepareTestFile prepares one file for testing.
func prepareTestFile() (*os.File, error) {

	file, err := ioutil.TempFile("", "load_test_*.properties")
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString("key1 = value1\r\n# comment...\nkey2= value2\nkey3=value3\r\n")
	if err != nil {
		return nil, err
	}

	file.Seek(0, io.SeekStart)
	return file, nil
}

// checkProperties checks this properties and returns an error if failed.
func checkProperties(properties *Properties) error {

	data := properties.data
	if len(data) != 3 {
		return fmt.Errorf("len of data (%d) is wrong", len(data))
	}

	for i := 1; i <= 3; i++ {
		index := strconv.Itoa(i)
		if value, ok := data["key" + index]; !ok || value != "value" + index {
			return fmt.Errorf("value of key%s (%s) is wrong", index, value)
		}
	}
	return nil
}

// go test -v -cover -run=^TestLoad$
func TestLoad(t *testing.T) {

	file, err := prepareTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	fileName := file.Name()
	t.Log("test file is", fileName)
	properties, err := Load(fileName)
	if err != nil {
		t.Fatal(err)
	}

	if err = checkProperties(properties); err != nil {
		t.Fatal(err)
	}
}

// go test -v -cover -run=^TestLoadFrom$
func TestLoadFrom(t *testing.T) {

	file, err := prepareTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	t.Log("test file is", file.Name())
	properties, err := LoadFrom(file)
	if err != nil {
		t.Fatal(err)
	}

	if err = checkProperties(properties); err != nil {
		t.Fatal(err)
	}
}
