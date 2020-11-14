// Author: fish
// Email: yezijie@bilibili.com
// Created at 2020/11/12 18:01

package props

import (
	"bytes"
	"io"
	"io/ioutil"
)

var (
	cr = []byte{'\r'}
	lf = []byte{'\n'}

	keyValueSeparator = []byte{'='}
	commentPrefix = []byte{'#'}
)

func Load(file string) (*Properties, error) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return parseBytes(data), nil
}

func LoadFrom(reader io.Reader) (*Properties, error) {

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return parseBytes(data), nil
}

func parseBytes(data []byte) *Properties {

	properties := newProperties()
	lines := bytes.Split(data, lf)
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if !isComment(line) {
			filledWithKv(properties, kvOf(line))
		}
	}
	return properties
}

func isComment(line []byte) bool {
	return bytes.HasPrefix(line, commentPrefix)
}

func kvOf(line []byte) [][]byte {
	line = bytes.TrimSuffix(line, cr)
	return bytes.SplitN(line, keyValueSeparator, 2)
}

func filledWithKv(properties *Properties, kv [][]byte) {
	if len(kv) == 2 {
		key := bytes.TrimSpace(kv[0])
		value := bytes.TrimSpace(kv[1])
		properties.data[string(key)] = string(value)
	}
}
