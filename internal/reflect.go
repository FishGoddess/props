package internal

import (
	"fmt"
	"reflect"
)

const (
	tagKey = "props"
)

// getTag gets tag from field with tagKey.
func getTag(field *reflect.StructField) (string, bool) {
	tag, ok := field.Tag.Lookup(tagKey)
	if !ok {
		tag = field.Name
	}

	if tag == "-" {
		return "", false
	}

	return tag, true
}

// mapToStruct maps m to v.
func mapToStruct(entries map[string][]string, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("props: the kind of v is %s, not struct", val.Kind())
	}

	return nil // TODO
}

// mapFromStruct maps v to m.
func mapFromStruct(v interface{}) (entries map[string][]string, err error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("props: the kind of v is %s, not struct", val.Kind())
	}

	entries = make(map[string][]string, 16)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		//tag := getTag(&field)

		switch field.Kind() {
		case reflect.Struct:
		case reflect.Array:
		case reflect.Slice:
		case reflect.Map:
		}
	}

	return entries, nil
}
