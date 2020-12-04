package passport

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const tag = "passport"

// Unmarshal a passport into a struct.
func Unmarshal(data []byte, v interface{}) error {
	data = append(data, ' ')
	fields := map[string]string{}
	for i := 0; i < len(data); {
		n, key, err := field(data[i:], ":")
		if err != nil {
			break
		}
		i += n + 1
		n, value, err := field(data[i:], " \n")
		if err != nil {
			break
		}
		i += n + 1
		fields[string(key)] = string(value)
	}

	s := reflect.ValueOf(v).Elem()
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tag := t.Field(i).Tag.Get(tag)
		var required bool
		if strings.HasPrefix(tag, "required,") {
			tag = strings.Replace(tag, "required,", "", 1)
			required = true
		}
		value, ok := fields[tag]
		if !ok && required {
			return fmt.Errorf("missing field %q", tag)
		}
		if !ok {
			continue
		}
		switch f.Kind() {
		case reflect.Int:
			v, err := strconv.ParseInt(string(value), 10, 64)
			if err != nil {
				return fmt.Errorf("could not parse %q (value: %q): %v", tag, string(value), err)
			}
			f.SetInt(v)
		case reflect.String:
			f.SetString(string(value))
		}
	}

	return nil
}

func field(data []byte, delim string) (int, []byte, error) {
	j := bytes.IndexAny(data, delim)
	if j < 0 {
		return 0, nil, fmt.Errorf("end of input")
	}
	return j, data[:j], nil
}
