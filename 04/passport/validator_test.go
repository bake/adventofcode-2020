package passport_test

import (
	"testing"

	"github.com/bake/adventofcode-2020/04/passport"
)

func TestValidate(t *testing.T) {
	tt := []struct {
		v     interface{}
		valid bool
	}{
		{
			v: struct {
				String string `validate:"min=3,max=6"`
			}{"fo"},
			valid: false,
		},
		{
			v: struct {
				String string `validate:"min=3,max=6"`
			}{"foo"},
			valid: true,
		},
		{
			v: struct {
				String string `validate:"min=3,max=6"`
			}{"foobarbaz"},
			valid: false,
		},
		{
			v: struct {
				Number int `validate:"min=3,max=6"`
			}{3},
			valid: true,
		},
		{
			v: struct {
				Number int `validate:"min=3,max=6"`
			}{42},
			valid: false,
		},
	}
	for i, tc := range tt {
		err := passport.Validate(tc.v)
		if tc.valid && err != nil {
			t.Errorf("expected validation %d to be %v, got %v (%v)", i+1, tc.valid, err == nil, err)
		}
	}
}
