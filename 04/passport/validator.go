package passport

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const validateTag = "validate"

// Validate a struct.
func Validate(v interface{}) error {
	s := reflect.ValueOf(v)
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		raw := t.Field(i).Tag.Get(validateTag)
		tags := strings.Split(raw, ",")
		for _, t := range tags {
			if len(t) == 0 {
				continue
			}
			val, err := parseValidator(t)
			if err != nil {
				return fmt.Errorf("could not parse %q: %v", t, err)
			}
			f := s.Field(i)
			if !val.validate(f.Interface()) {
				return fmt.Errorf("%v does not satisfy %q", f, t)
			}
		}
	}
	return nil
}

type validator interface {
	validate(v interface{}) bool
}

func parseValidator(raw string) (validator, error) {
	parts := strings.SplitN(raw, "=", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("value missing")
	}
	switch parts[0] {
	case "min":
		min, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		return minValidator{min: min}, nil
	case "max":
		max, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		return maxValidator{max: max}, nil
	default:
		return nil, fmt.Errorf("validator %q does not exist", parts[0])
	}
}

type minValidator struct {
	min int
}

func (val minValidator) validate(v interface{}) bool {
	switch v.(type) {
	case int:
		return v.(int) >= val.min
	case string:
		return len(v.(string)) >= val.min
	default:
		return false
	}
}

type maxValidator struct {
	max int
}

func (val maxValidator) validate(v interface{}) bool {
	switch v.(type) {
	case int:
		return v.(int) <= val.max
	case string:
		return len(v.(string)) <= val.max
	default:
		return false
	}
}
