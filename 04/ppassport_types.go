package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type height struct {
	value int
	unit  string
}

func (h height) String() string { return fmt.Sprintf("%d%s", h.value, h.unit) }

func (h *height) UnmarshalPassport(v string) error {
	parts := regexp.MustCompile(`^([\d]+)(in|cm)$`).FindStringSubmatch(v)
	if len(parts) != 3 {
		return fmt.Errorf("invalid syntax %q", v)
	}
	var err error
	h.value, err = strconv.Atoi(parts[1])
	h.unit = parts[2]
	switch h.unit {
	case "cm":
		if 150 > h.value || h.value > 193 {
			return fmt.Errorf("invalid size: %s", h)
		}
	case "in":
		if 59 > h.value || h.value > 76 {
			return fmt.Errorf("invalid size: %s", h)
		}
	}
	return err
}

type color struct{ value string }

func (c color) String() string { return fmt.Sprintf("#%s", c.value) }

func (c *color) UnmarshalPassport(v string) error {
	if !regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(v) {
		return fmt.Errorf("invalid syntax %q", v)
	}
	c.value = v[1:]
	return nil
}

type eyeColor struct{ value string }

func (c eyeColor) String() string { return c.value }

func (c *eyeColor) UnmarshalPassport(v string) error {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range colors {
		if v == color {
			c.value = v
			return nil
		}
	}
	return fmt.Errorf("unexpected color %q", v)
}
