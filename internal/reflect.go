package internal

import (
	"fmt"
	"reflect"
)

const (
	tagKey = "props"
)

// mapToStruct maps m to v.
func mapToStruct(entries map[string][]string, v interface{}) error {
	return nil // TODO
}

// mapFromStruct maps v to m.
func mapFromStruct(v interface{}) (entries map[string][]string, err error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("props: the kind of v is %s, not struct", t.Kind())
	}

	entries = make(map[string][]string, 16)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup(tagKey)
		if !ok {
			tag = field.Name
		}

		fieldType := field.Type
		fieldType.Kind() // TODO
		_ = tag
	}

	return entries, nil
}
