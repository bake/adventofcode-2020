package passport_test

import (
	"testing"

	"github.com/bake/adventofcode-2020/04/passport"
)

func TestUnmarshal(t *testing.T) {
	tt := []struct {
		data  []byte
		valid bool
	}{
		{
			data:  []byte("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm"),
			valid: true,
		},
		{
			data:  []byte("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929"),
			valid: false,
		},
		{
			data:  []byte("hcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm"),
			valid: true,
		},
		{
			data:  []byte("hcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in"),
			valid: false,
		},
	}
	type port struct {
		BirthYear      int    `passport:"required,byr"`
		IssueYear      int    `passport:"required,iyr"`
		ExpirationYear int    `passport:"required,eyr"`
		Height         string `passport:"required,hgt"`
		HairColor      string `passport:"required,hcl"`
		EyeColor       string `passport:"required,ecl"`
		PassportID     int    `passport:"required,pid"`
		CountryID      int    `passport:"cid"`
	}
	for i, tc := range tt {
		var p port
		err := passport.Unmarshal(tc.data, &p)
		if err != nil && tc.valid {
			t.Errorf("expected passport %d to be %v, got %v (%v)", i+1, tc.valid, err == nil, err)
		}
	}
}
