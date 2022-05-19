package common

import (
	"errors"
	"strconv"
)

var unitMap = map[string]struct{}{
	"d": {},
	"w": {},
	"m": {},
	"y": {},
}

type LongDuration struct {
	Days   int
	Months int
	Years  int
}

// ParseLongDuration parses a duration condition string.
// A duration string is a sequence of
// whole number and a unit suffix,
// such as "1m", "1d" or "2y15w".
// Valid time units are "d", "w", "m", "y".
func ParseLongDuration(s string) (*LongDuration, error) {

	// ([0-9]*(\.[0-9]*)?[a-z]+)+
	orig := s
	checkedMap := make(map[string]int)

	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return &LongDuration{}, nil
	}
	if s == "" {
		return nil, errors.New("invalid duration `" + orig + "`")
	}

	for s != "" {
		var err error
		var value int
		var unit string

		c := s[0]
		// The next character must be [0-9]
		if c < '0' || c > '9' {
			return nil, errors.New("invalid duration `" + orig + "`")
		}

		// Consume [0-9]+(\.[0-9]*)?
		value, s, err = leadingInt(s)
		if err != nil {
			return nil, errors.New("invalid duration `" + orig + "`")
		}

		// Consume unit
		unit, s = leadingLetters(s)
		if unit == "" {
			return nil, errors.New("missing unit in duration `" + orig + "`")
		}
		_, ok := unitMap[unit]
		if !ok {
			return nil, errors.New("unknown unit `" + unit + "` in duration `" + orig + "`")
		}
		checkedMap[unit] = value
	}
	return &LongDuration{
		Days:   checkedMap["d"] + checkedMap["w"]*7,
		Months: checkedMap["m"],
		Years:  checkedMap["y"],
	}, nil
}

// leadingInt consumes the leading [0-9]* from s.
func leadingInt(s string) (x int, rem string, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
	}

	x, err = strconv.Atoi(s[:i])
	if err != nil {
		return 0, "", err
	}

	return x, s[i:], nil
}

// leadingLetters consumes the leading [a-z]* from s.
func leadingLetters(s string) (x string, rem string) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < 'a' || c > 'z' {
			break
		}
	}

	return s[:i], s[i:]
}
