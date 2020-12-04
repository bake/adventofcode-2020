package passport_test

import (
	"testing"

	"github.com/bake/adventofcode-2020/04/passport"
)

func TestValidate(t *testing.T) {
	v := struct {
		String string `validate:"min=3,max=6"`
	}{
		String: "foo",
	}
	if err := passport.Validate(v); err != nil {
		t.Fatal(err)
	}
	t.Fatal()
}
