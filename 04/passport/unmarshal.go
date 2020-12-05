package passport

import (
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
	fs := map[string]string{}
	s := NewScanner(data)
	for s.Next() {
		fs[string(s.pair.key)] = string(s.pair.value)
	}
	if err := s.Error(); err != nil {
		return fmt.Errorf("could not parse passport: %v", err)
	}

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
