package passport

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

const unmarshalTag = "passport"

// The Unmarshaller interface allows to decode into custom types.
type Unmarshaller interface {
	UnmarshalPassport(string) error
}

// Unmarshal a passport into a struct.
func Unmarshal(data []byte, v interface{}) error {
	fs := fields(data)
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("v must be a pointer to a struct")
	}
	ts := reflect.TypeOf(v).Elem()
	if ts.Kind() != reflect.Struct {
		return fmt.Errorf("v must be a pointer to a struct")
	}
	vs := reflect.ValueOf(v).Elem()

	for i := 0; i < ts.NumField(); i++ {
		if !vs.Field(i).CanSet() {
			continue
		}
		tag, ok := ts.Field(i).Tag.Lookup(unmarshalTag)
		if !ok {
			continue
		}
		value, ok := fs[tag]
		if !ok {
			continue
		}
		if err := set(ts.Field(i).Type.Kind(), vs.Field(i), value); err != nil {
			return fmt.Errorf("could not set %q: %v", ts.Field(i).Name, err)
		}
	}
	return nil
}

func fields(data []byte) map[string]string {
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
	return fields
}

func field(data []byte, delim string) (int, []byte, error) {
	j := bytes.IndexAny(data, delim)
	if j < 0 {
		return 0, nil, fmt.Errorf("end of input")
	}
	return j, data[:j], nil
}

func set(kind reflect.Kind, v reflect.Value, data string) error {
	switch kind {
	case reflect.Int:
		val, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return fmt.Errorf("could not parse %q: %v", kind, err)
		}
		v.SetInt(val)
		return nil
	case reflect.String:
		v.SetString(data)
		return nil
	default:
		if g, ok := v.Addr().Interface().(Unmarshaller); ok {
			if err := g.UnmarshalPassport(data); err != nil {
				return fmt.Errorf("could not unmarshal %q: %v", kind, err)
			}
			return nil
		}
		return fmt.Errorf("unexpected type %q", kind)
	}
}
